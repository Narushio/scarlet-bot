package handler

import (
	"fmt"

	"github.com/Narushio/scarlet-bot/thirdparty/botgo/dto"
	"github.com/Narushio/scarlet-bot/thirdparty/botgo/event"
)

// ThreadEvent 论坛主贴事件
func ThreadEvent() event.ThreadEventHandler {
	return func(event *dto.WSPayload, data *dto.WSThreadData) error {
		fmt.Println(event, data)
		return nil
	}
}
