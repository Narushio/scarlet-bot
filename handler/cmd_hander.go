package handler

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

type CmdHandler interface {
	buildMessageToCreate(cmd message.CMD) *dto.MessageToCreate
}
