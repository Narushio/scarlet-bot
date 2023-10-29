package processor

import (
	"context"

	"github.com/Narushio/scarlet-bot/plugin/qqx5"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

type CmdHandler struct {
	Processor     *Processor
	MessageData   *dto.WSMessageData
	ATMessageData *dto.WSATMessageData
}

func NewCmdHandler(p *Processor, data *dto.WSMessageData, atData *dto.WSATMessageData) *CmdHandler {
	return &CmdHandler{
		Processor:     p,
		MessageData:   data,
		ATMessageData: atData,
	}
}

func (c *CmdHandler) HandleCmd(ctx context.Context, cmd *message.CMD) error {
	switch cmd.Cmd {
	case "/qqx5":
		q := qqx5.New(c.Processor)
		err := q.SendReplay(ctx, cmd.Content, c.MessageData, c.ATMessageData)
		if err != nil {
			return err
		}
	}
	return nil
}
