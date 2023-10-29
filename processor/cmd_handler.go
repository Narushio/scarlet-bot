package processor

import (
	"context"

	"github.com/Narushio/scarlet-bot/api"
	"github.com/Narushio/scarlet-bot/plugin/qqx5"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

type CmdHandler struct {
	TencentAPI    *api.TencentAPI
	MessageData   *dto.WSMessageData
	ATMessageData *dto.WSATMessageData
}

func NewCmdHandler(api *api.TencentAPI, data *dto.WSMessageData, atData *dto.WSATMessageData) *CmdHandler {
	return &CmdHandler{
		TencentAPI:    api,
		MessageData:   data,
		ATMessageData: atData,
	}
}

func (c *CmdHandler) HandleCmd(ctx context.Context, cmd *message.CMD) error {
	switch cmd.Cmd {
	case "/qqx5":
		q := qqx5.New(c.TencentAPI)
		err := q.SendReplay(ctx, cmd.Content, c.MessageData, c.ATMessageData)
		if err != nil {
			return err
		}
	}
	return nil
}
