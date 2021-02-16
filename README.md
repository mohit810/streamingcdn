# This server accepts video and audio via webrtc and converts them in Adaptive HLS
This project is in Progress(ui,cdn integration), `Webrtc -> HLS` part is Complete.

### Steps for running this Server (I recommend using [local-testing](https://github.com/mohit810/streamingcdn/tree/local-testing) branch)

#### Steps to follow for using this project
1) Install or upgrade Golang ver. 1.15 or above
2) Install or upgrade Ffmpeg ver. 4.3.1-5(available in `Debian Sid`) and windows ffmpeg v4.3
3) Download the project using Git command `git clone https://github.com/mohit810/streamingcdn` 
4) To start the server open the new folder created by Git and use the command `go run main.go` (Use terminal for running the server)

* Congratulations!! Now the server is up and running.
* As a starting point for testing you can use the [Web UI](https://github.com/mohit810/streamingcdn-web-ui)

### Current State
1) Users can connect via `POST` request and start streaming(1.5Mbps speed is hardcoded for the stream).
2) The Server recieves the stream and converts it into HLS.
3) [Native Android-UI](https://github.com/mohit810/android-webrtc)
4) [Web UI](https://github.com/mohit810/streamingcdn-web-ui)

### Currently in Development
1) Integrating FFmpeg.(Completed for now, further optimization will be done later)
2) Allowing multiple users to connect and start broadcasting.(not in priority, at this stage.)
3) CDN Integration ( Post UI development)
4) Front-end UI Development in ios. 

### Final Goal
To receive the broadcast, Convert it into hls in various quality then push those to cdn to serve to the User.

Before using this solution you should set-up pion/webrtc/v3 ([Go Modules](https://blog.golang.org/using-go-modules) are mandatory for using Pion WebRTC. So make sure you set export GO111MODULE=on, and explicitly specify /v3 when importing.).

### POST API used by the broadcaster for connecting to the server

`{
    "sdp":"",
    "streamKey":""
}`

* Playlists and all the `.ts` files are served at `/watch/streamkey/` (here streamKey is the string that the broadcaster sends in the `POST` request when connecting to the server)

## Big Thanks to the following 

* [Sean Der](https://github.com/Sean-Der) at [Poin/webrtc](https://github.com/pion/webrtc)
* [Harrison](https://github.com/grantfayvor)
* [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
