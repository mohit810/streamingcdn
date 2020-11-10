let divSelectRoom = document.getElementById("selectRoom")
let inputstreamKey = document.getElementById("streamkey")
let signalingContainer = document.getElementById('signalingContainer')
let createSessionButton = document.getElementsByClassName('createSessionButton')
let remoteSessionDescription = document.getElementById('remoteSessionDescription')
let localSessionDescription = document.getElementById('localSessionDescription')
let video1 = document.getElementById('video1')

let streamKey, encryptedSdp, PublisherFlag, uid

/* eslint-env browser */
var log = msg => {
  document.getElementById('logs').innerHTML += msg + '<br>'
}

const hdConstraints = {
  audio: true,
  video: {
    width: { max: 1920, ideal: 1280 },
    height: { max: 1080, ideal: 720 }
  }
};

let displayVideo = video => {
    var el = document.createElement('video')
    el.srcObject = video
    el.autoplay = true
    el.muted = true
    el.width = 160
    el.height = 120

    document.getElementById('localVideos').appendChild(el)
    return video
}

function postRequest () {
    var data = JSON.stringify({
        "sdp": encryptedSdp,
        "streamKey": streamKey
    })
    console.log(data);
    const url = `${window.location.origin}/sdp`;
    (async () => {
        const rawResponse = await fetch(url, {
        method : "POST",
        //body: new FormData(document.getElementById("inputform")),
        // -- or --
        body : data,
        headers :{
            'Content-Type': 'application/json'
        }
    });
    const content = await rawResponse.json();
        remoteSessionDescription.value = content.sdp
        window.startSession()
})();
}
window.createSession = isPublisher => {
  PublisherFlag = isPublisher
  if (inputstreamKey.value === '') {
    alert("please enter a room name.")
  } else{
    streamKey = inputstreamKey.value
    let pc = new RTCPeerConnection({
      iceServers: [
        {'urls': 'stun:stun.services.mozilla.com'},
        {'urls': 'stun:stun.l.google.com.19302'}
      ]
    })
  pc.oniceconnectionstatechange = e => log(pc.iceConnectionState)
  pc.onicecandidate = event => {
    if (event.candidate === null) {
      encryptedSdp = btoa(JSON.stringify(pc.localDescription))
      localSessionDescription.value = encryptedSdp
        postRequest();
    }
  }

  if (isPublisher) {
      navigator.mediaDevices.getUserMedia(hdConstraints)
          .then(stream => {
              stream.getTracks().forEach(function(track) {
                  pc.addTrack(track, stream);
              });
              displayVideo(stream);
            pc.createOffer()
                .then(d => {
                  pc.setLocalDescription(d)
                }).catch(log)
          }).catch(log)
  }else{
      pc.addTransceiver('audio', {'direction': 'recvonly'})
      pc.addTransceiver('video', {'direction': 'recvonly'})
      pc.createOffer()
          .then(d => pc.setLocalDescription(d))
          .catch(log)

      pc.ontrack = function (event) {
        var el = video1
        el.srcObject = event.streams[0]
        el.autoplay = true
        el.controls = true
      }
  }
  window.startSession = () => {
        let sd = remoteSessionDescription.value
        if (sd === '') {
          return alert('Session Description must not be empty')
        }
        try {
          pc.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(sd))))
        } catch (e) {
          alert(e)
        }
      }

  let btns = createSessionButton
  for (let i = 0; i < btns.length; i++) {
    btns[i].style = 'display: none'
  }
  divSelectRoom.style = "display: none"
  signalingContainer.style = 'display: block'
}
}
