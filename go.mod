module SystemgeSampleProxy

go 1.23

//replace github.com/neutralusername/systemge => ../Systemge

require github.com/neutralusername/systemge v0.0.0-20241024145243-e1a849f9772a

require (
	github.com/gorilla/websocket v1.5.3 // indirect
	golang.org/x/oauth2 v0.21.0 // indirect
)
