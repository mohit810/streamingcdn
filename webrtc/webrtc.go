package webrtc

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os/exec"
	"time"

	"github.com/mohit810/streamingcdn/encryptor"
	"github.com/pion/rtcp"
	"github.com/pion/webrtc/v3"
)

type udpConn struct {
	conn *net.UDPConn
	port int
}

// CreateWebRTCConnection function
func CreateWebRTCConnection(offerStr string) (answer webrtc.SessionDescription, err error) {

	defer func() {
		if e, ok := recover().(error); ok {
			err = e
		}
	}()

	// Create a MediaEngine object to configure the supported codec
	m := webrtc.MediaEngine{}

	// Setup the codecs you want to use.
	m.RegisterCodec(webrtc.NewRTPOpusCodec(webrtc.DefaultPayloadTypeOpus, 48000))
	m.RegisterCodec(webrtc.NewRTPH264Codec(webrtc.DefaultPayloadTypeH264, 90000))

	// Create the API object with the MediaEngine
	api := webrtc.NewAPI(webrtc.WithMediaEngine(m))

	// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := api.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// Allow us to receive 1 audio track, and 1 video track
	if _, err = peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio); err != nil {
		panic(err)
	} else if _, err = peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo); err != nil {
		panic(err)
	}

	go func(peerConnection *webrtc.PeerConnection) {
		// Create context
		ctx, cancel := context.WithCancel(context.Background())

		// Create a local addr
		var laddr *net.UDPAddr
		if laddr, err = net.ResolveUDPAddr("udp", "127.0.0.1:"); err != nil {
			fmt.Println(err)
			cancel()
		}

		// Prepare udp conns
		udpConns := map[string]*udpConn{
			"audio": {port: 4000},
			"video": {port: 4002},
		}
		for _, c := range udpConns {
			// Create remote addr
			var raddr *net.UDPAddr
			if raddr, err = net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", c.port)); err != nil {
				fmt.Println(err)
				cancel()
			}

			// Dial udp
			if c.conn, err = net.DialUDP("udp", laddr, raddr); err != nil {
				fmt.Println(err)
				cancel()
			}
			defer func(conn net.PacketConn) {
				if closeErr := conn.Close(); closeErr != nil {
					fmt.Println(closeErr)
				}
			}(c.conn)
		}

		startFFmpeg(ctx)

		// Set a handler for when a new remote track starts, this handler will forward data to
		// our UDP listeners.
		// In your application this is where you would handle/process audio/video
		peerConnection.OnTrack(func(track *webrtc.Track, receiver *webrtc.RTPReceiver) {
			fmt.Println("on track called")

			// Retrieve udp connection
			c, ok := udpConns[track.Kind().String()]
			if !ok {
				return
			}

			// Send a PLI on an interval so that the publisher is pushing a keyframe every rtcpPLIInterval
			go func() {
				ticker := time.NewTicker(time.Second * 2)
				for range ticker.C {
					if rtcpErr := peerConnection.WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: track.SSRC()}}); rtcpErr != nil {
						fmt.Println(rtcpErr)
					}
					if rtcpSendErr := peerConnection.WriteRTCP([]rtcp.Packet{&rtcp.ReceiverEstimatedMaximumBitrate{Bitrate: 1500000, SenderSSRC: track.SSRC()}}); rtcpSendErr != nil {
						fmt.Println(rtcpSendErr)
					}
				}
			}()

			b := make([]byte, 1500)
			for {
				// Read
				n, readErr := track.Read(b)
				if readErr != nil {
					fmt.Println(readErr)
				}

				// Write
				if _, err = c.conn.Write(b[:n]); err != nil {
					fmt.Println(err)
				}
			}
		})

		// in a production application you should exchange ICE Candidates via OnICECandidate
		peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
			fmt.Println(candidate)
		})

		// Set the handler for ICE connection state
		// This will notify you when the peer has connected/disconnected
		peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
			fmt.Printf("Connection State has changed %s \n", connectionState.String())

			if connectionState == webrtc.ICEConnectionStateConnected {
				fmt.Println("ICE connection was successful")
			} else if connectionState == webrtc.ICEConnectionStateFailed ||
				connectionState == webrtc.ICEConnectionStateDisconnected {
				cancel()
			}
		})

		// Wait for context to be done
		<-ctx.Done()
		peerConnection.Close()

	}(peerConnection)

	// Wait for the offer to be pasted
	offer := webrtc.SessionDescription{}
	encryptor.Decode(offerStr, &offer)
	// Set the remote SessionDescription
	if err = peerConnection.SetRemoteDescription(offer); err != nil {
		panic(err)
	}

	// Create answer
	answer, err = peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	// Sets the LocalDescription, and starts our UDP listeners
	if err = peerConnection.SetLocalDescription(answer); err != nil {
		panic(err)
	}

	return
}

func startFFmpeg(ctx context.Context) {
	// Create a ffmpeg process that consumes MKV via stdin, and broadcasts out to Stream URL
	//CURRENTLY THIS IS FFMPEG PROCESS RUNS ON CPU NOT ON GPU
	ffmpeg := exec.CommandContext(ctx, "ffmpeg", "-protocol_whitelist", "file,udp,rtp", "-i", "rtp-forwarder.sdp", "-map", "0:v:0", "-map", "0:a:0", "-map", "0:v:0", "-map", "0:a:0", "-map", "0:v:0", "-map", "0:a:0", "-map", "0:v:0", "-map", "0:a:0", "-c:v", "h264", "-profile:v", "main", "-pix_fmt", "yuv420p", "-preset", "faster", "-crf", "20", "-sc_threshold", "0", "-g", "48", "-keyint_min", "48", "-c:a", "aac", "-ac", "2", "-ar", "48000", "-filter:v:0", "scale=w=640:h=360:force_original_aspect_ratio=decrease", "-maxrate:v:0", "856k", "-bufsize:v:0", "1200k", "-b:v:0", "800k", "-b:a:0", "96k", "-filter:v:1", "scale=w=842:h=480:force_original_aspect_ratio=decrease", "-maxrate:v:1", "1498k", "-bufsize:v:1", "2100k", "-b:v:1", "1400k", "-b:a:1", "128k", "-filter:v:2", "scale=w=1280:h=720:force_original_aspect_ratio=decrease", "-maxrate:v:2", "2996k", "-bufsize:v:2", "4200k", "-b:v:2", "2800k", "-b:a:2", "128k", "-filter:v:3", "scale=w=1920:h=1080:force_original_aspect_ratio=decrease", "-maxrate:v:3", "5350k", "-bufsize:v:3", "7500k", "-b:v:3", "5000k", "-b:a:3", "192k", "-var_stream_map", "v:0,a:0 v:1,a:1 v:2,a:2 v:3,a:3", "-hls_time", "4", "-hls_playlist_type", "vod", "-master_pl_name", "master.m3u8", "-hls_segment_filename", "video_%v_%03d.ts", "quality_%v.m3u8") //not creating master and different resolution playlist
	ffmpegOut, _ := ffmpeg.StderrPipe()
	if err := ffmpeg.Start(); err != nil {
		panic(err)
	}

	go func() {
		scanner := bufio.NewScanner(ffmpegOut)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	/*same ffmpeg command for terminal with static input*/
	//ffmpeg -hide_banner -re -i sample720.mp4 \
	//-map 0:v:0 -map 0:a:0 -map 0:v:0 -map 0:a:0 -map 0:v:0 -map 0:a:0 -map 0:v:0 -map 0:a:0 \
	//-c:v h264 -profile:v main -pix_fmt yuv420p -preset faster -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -c:a aac -ac 2 -ar 48000 \
	//-filter:v:0 scale=w=640:h=360:force_original_aspect_ratio=decrease  -maxrate:v:0 856k  -bufsize:v:0 1200k -b:v:0 800k -b:a:0 96k \
	//-filter:v:1 scale=w=842:h=480:force_original_aspect_ratio=decrease  -maxrate:v:1 1498k -bufsize:v:1 2100k -b:v:1 1400k -b:a:1 128k \
	//-filter:v:2 scale=w=1280:h=720:force_original_aspect_ratio=decrease -maxrate:v:2 2996k -bufsize:v:2 4200k -b:v:2 2800k -b:a:2 128k \
	//-filter:v:3 scale=w=1920:h=1080:force_original_aspect_ratio=decrease -maxrate:v:3 5350k -bufsize:v:3 7500k -b:v:3 5000k -b:a:3 192k \
	//-var_stream_map "v:0,a:0 v:1,a:1 v:2,a:2 v:3,a:3" -hls_time 4 -hls_playlist_type vod \
	//-master_pl_name master.m3u8 -hls_segment_filename video_%v_%03d.ts quality_%v.m3u8
}
