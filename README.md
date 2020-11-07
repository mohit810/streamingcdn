# Signalling Pion Webrtc server via `POST` request

This project is in Progress.

### Current State
1) Users can connect via `POST` request and start streaming(1.5Mbps speed is hardcoded for the stream).
2) The Server recieves the stream and saves it with the unique Id as the name.

### Currently in Devlopment
1) Integrating FFmpeg.
2) Allowing multiple users to connect start broadcasting.

### Final Goal
To receive the broadcast, Convert it into hls in various quality then push those to cdn to serve to the User.

Before using this solution you should set-up have pion/webrtc/v3 ([Go Modules](https://blog.golang.org/using-go-modules) are mandatory for using Pion WebRTC. So make sure you set export GO111MODULE=on, and explicitly specify /v3 when importing.).

### Open broadcast example page
[localhost:8000](http://localhost:8000/) You should see two buttons 'Publish a Broadcast' . 

### Run Application
#### Linux/macOS/windows
Run `main.go`

### Start a publisher

* Click `Publish a Broadcast`
* For Communicating with server you have to request the server via `POST` method and paste the sdp obtained from the browser as well as uid. The `application` will respond with an offer as a response to the `POST`, paste this into the second input field. Then press `Start Session`.

![](https://github.com/mohit810/streamingcdn/blob/master/Screenshot.png)

## Big Thanks to the following 

* [Sean Der](https://github.com/Sean-Der) at [Poin/webrtc](https://github.com/pion/webrtc)
* [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
