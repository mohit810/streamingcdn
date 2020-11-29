# This server accepts video and audio via webrtc and converts them in Adaptive HLS
This project is in Progress(ui,cdn integration), `Webrtc -> HLS` part is Complete.

### Run Application

#### Steps to follow for using this project
1) Install or upgrade Golang ver. 1.15 or above
2) Install or upgrade Ffmpeg ver. 4.3.1-5 or above (Currently available in `Debian Sid`)
3) Download the project using Git command `git clone https://github.com/mohit810/streamingcdn` 
4) To start the server use the command `go run main.go` (Use terminal for running the server)

* Congratulations!! Now the server is up and running.
* As a starting point for testing you can use the [Web UI](https://github.com/mohit810/streamingcdn-web-ui)

### Current State
1) Users can connect via `POST` request and start streaming(1.5Mbps speed is hardcoded for the stream).
2) The Server recieves the stream and converts it into HLS.

### Currently in Devlopment
1) Integrating FFmpeg.(Completed for now, further optimization will be done later)
2) Allowing multiple users to connect and start broadcasting.(not in priority, at this stage.)
3) CDN Integration ( Post UI development)
4) Front-end UI Development in android, web, ios. (Current priority is android)

### Final Goal
To receive the broadcast, Convert it into hls in various quality then push those to cdn to serve to the User.

Before using this solution you should set-up pion/webrtc/v3 ([Go Modules](https://blog.golang.org/using-go-modules) are mandatory for using Pion WebRTC. So make sure you set export GO111MODULE=on, and explicitly specify /v3 when importing.).

### POST API used by the broadcaster for connecting to the server

`{
    "sdp":"",
    "streamKey":""
}`


## Big Thanks to the following 

* [Sean Der](https://github.com/Sean-Der) at [Poin/webrtc](https://github.com/pion/webrtc)
* [Harrison](https://github.com/grantfayvor)
* [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
