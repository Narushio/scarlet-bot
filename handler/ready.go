package handler

import (
	"log"

	"github.com/Narushio/scarlet-bot/pkg/dto"
	"github.com/Narushio/scarlet-bot/pkg/event"
)

// Ready 自定义 Ready 感知连接成功事件
func Ready() event.ReadyHandler {
	return func(event *dto.WSPayload, data *dto.WSReadyData) {
		log.Println("ready event receive: ", data)
	}
}
