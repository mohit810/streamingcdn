# Signalling Pion Webrtc server via `POST` request

This project is in Progress.

### Software Req.
1) Golang >= 1.15
2) Ffmpeg >= 4.3.1-5 (Currently available in `Debian Sid`)

### Current State
1) Users can connect via `POST` request and start streaming(1.5Mbps speed is hardcoded for the stream).
2) The Server recieves the stream and converts it into HLS (give the location where you want the files to be saved([code for this](https://github.com/mohit810/streamingcdn/blob/8125fb4afa63398dabc1048bc98e2d4266fa4dcf/ffmpeg/ffmpeg.go#L13)) ) .
3) The Web UI is complete 

### Currently in Devlopment
1) Integrating FFmpeg.(Completed for now, further optimization will be done later)
2) Allowing multiple users to connect and start broadcasting.(not in priority, at this stage.)
3) CDN Integration ( Post UI development)
4) Front-end UI Development in android, web, ios. (Current priority is android)
### Final Goal
To receive the broadcast, Convert it into hls in various quality then push those to cdn to serve to the User.

Before using this solution you should set-up pion/webrtc/v3 ([Go Modules](https://blog.golang.org/using-go-modules) are mandatory for using Pion WebRTC. So make sure you set export GO111MODULE=on, and explicitly specify /v3 when importing.).

### Web ui page
[Web UI](https://github.com/mohit810/streamingcdn-web-ui) 

### Run Application
#### Linux/macOS/windows
Run `go run main.go`

### Start a publisher

* Paste any random thing u wanna past.
* Click `Publish a Broadcast` and now you don't have to do anything.
* Communicating with server is done by the js itself.  
* If you want to know the request that is sent to server via `POST` method, refer to the screenshot attached below. The `application` will respond with an offer as a response to the `POST`.

![](https://github.com/mohit810/streamingcdn/blob/main/Screenshot.png)

## Big Thanks to the following 

* [Sean Der](https://github.com/Sean-Der) at [Poin/webrtc](https://github.com/pion/webrtc)
* [Harrison](https://github.com/grantfayvor)
* [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
