package main

import (
	"SystemgeSampleProxy/proxyServer"
	"time"
)

func main() {
	server, err := proxyServer.New()
	if err != nil {
		panic(err)
	}
	if err := server.Start(); err != nil {
		panic(err)
	}
	<-make(<-chan time.Time)
}
