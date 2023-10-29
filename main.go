package main

import (
	"github.com/Narushio/scarlet-bot/browser"
)

func main() {
	browser.LaunchChromium()
	go setHttpServer(8000)
	launchBot()
}
