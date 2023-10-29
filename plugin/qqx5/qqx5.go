package qqx5

import (
	"context"

	"github.com/Narushio/scarlet-bot/processor"
	"github.com/tencent-connect/botgo/dto"
)

type Plugin struct {
	Processor *processor.Processor
}

func New(p *processor.Processor) *Plugin {
	return &Plugin{Processor: p}
}

func (p Plugin) SendReplay(ctx context.Context, content string, data *dto.WSMessageData, atData *dto.WSATMessageData) error {
	switch content {
	case "双排星动爆点":
		if err := p.sendDuoBurstTable(ctx, data, atData); err != nil {
			return err
		}
	}
	return nil
}

func (p Plugin) sendDuoBurstTable(ctx context.Context, data *dto.WSMessageData, atData *dto.WSATMessageData) error {
	//msgData := map[string]string{
	//	"content": "被动消息测试",
	//	"msg_id":  data.ID,
	//}
	//a, err := p.tencentAPI.ExtendedAPI.SendPicToChannelMsg(ctx, data.ChannelID, "resource/图片1.jpg", msgData)
	//if err != nil {
	//	log.Println(err)
	//}
	return nil
}
