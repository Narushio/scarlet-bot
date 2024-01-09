package handler

import (
	"fmt"

	"github.com/Narushio/scarlet-bot/thirdparty/botgo/dto"
	"github.com/Narushio/scarlet-bot/thirdparty/botgo/event"
)

// ChannelEvent 处理子频道事件
func ChannelEvent() event.ChannelEventHandler {
	return func(event *dto.WSPayload, data *dto.WSChannelData) error {
		fmt.Println(data)
		return nil
	}
}
