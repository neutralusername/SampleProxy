package main

import (
	"SystemgeSampleProxy/appProxy"
	"time"
)

func main() {
	appProxy.New()
	<-make(<-chan time.Time)
}
