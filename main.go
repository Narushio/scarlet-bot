package main

import (
	"log"

	"github.com/Narushio/scarlet-bot/bot"
	"github.com/Narushio/scarlet-bot/browser"
	"github.com/Narushio/scarlet-bot/plugin"
)

func main() {
	manager := plugin.NewManager()
	manager.Register()
	browser.LaunchChromium()
	go setHttpServer(8000)
	err := bot.NewScarletBot().LinkStart()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
