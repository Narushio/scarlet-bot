package handler

import (
	"fmt"

	"github.com/Narushio/scarlet-bot/thirdparty/botgo/dto"
	"github.com/Narushio/scarlet-bot/thirdparty/botgo/event"
)

// MemberEvent 处理成员变更事件
func MemberEvent() event.GuildMemberEventHandler {
	return func(event *dto.WSPayload, data *dto.WSGuildMemberData) error {
		fmt.Println(data)
		return nil
	}
}
