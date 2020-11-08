# Signalling Pion Webrtc server via `POST` request

This project is in Progress.

### Current State
1) Users can connect via `POST` request and start streaming(1.5Mbps speed is hardcoded for the stream).
2) The Server recieves the stream and forwards the rtp packets to twitch's rtmp link.

### Currently in Devlopment
1) Integrating FFmpeg.
2) Allowing multiple users to connect start broadcasting.

### Final Goal
To receive the broadcast, Convert it into hls in various quality then push those to cdn to serve to the User.

Before using this solution you should set-up pion/webrtc/v3 ([Go Modules](https://blog.golang.org/using-go-modules) are mandatory for using Pion WebRTC. So make sure you set export GO111MODULE=on, and explicitly specify /v3 when importing.).

### Open broadcast example page
[localhost:8000](http://localhost:8000/) 

### Run Application
#### Linux/macOS/windows
Run `main.go`

### Start a publisher

* Paste your Twitch stream Key
* Click `Publish a Broadcast` and now you don't have to do anything.
* Communicating with server is done by the js itself.  
* If you want to know the request that is sent to server via `POST` method, refer to the screenshot attached below. The `application` will respond with an offer as a response to the `POST`.

![](https://github.com/mohit810/streamingcdn/blob/dev-branch/Screenshot.png)

## Big Thanks to the following 

* [Sean Der](https://github.com/Sean-Der) at [Poin/webrtc](https://github.com/pion/webrtc)
* [Harrison](https://github.com/grantfayvor)
* [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
