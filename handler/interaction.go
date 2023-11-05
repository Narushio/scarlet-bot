package handler

import (
	"fmt"

	"github.com/Narushio/scarlet-bot/pkg/dto"
	"github.com/Narushio/scarlet-bot/pkg/event"
)

// Interaction 处理内联交互事件
func Interaction() event.InteractionEventHandler {
	return func(event *dto.WSPayload, data *dto.WSInteractionData) error {
		fmt.Println(data)
		//return proc.ProcessInlineSearch(data)
		return nil
	}
}
