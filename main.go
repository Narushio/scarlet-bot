package main

import (
	"log"

	"github.com/Narushio/scarlet-bot/bot"
	"github.com/Narushio/scarlet-bot/browser"
)

func main() {
	browser.LaunchChromium()
	go setHttpServer(8000)
	err := bot.NewScarletBot().LinkStart()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
