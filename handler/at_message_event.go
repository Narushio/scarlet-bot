package handler

import (
	"fmt"

	"github.com/Narushio/scarlet-bot/thirdparty/botgo/dto"
	"github.com/Narushio/scarlet-bot/thirdparty/botgo/event"
)

// ATMessageEvent 实现处理 at 消息的回调
func ATMessageEvent() event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		fmt.Println(data)
		return nil
	}
}
