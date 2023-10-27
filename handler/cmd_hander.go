package handler

import (
	"github.com/tencent-connect/botgo/dto/message"
)

type CmdHandler interface {
	SendReplay(cmd message.CMD)
}
