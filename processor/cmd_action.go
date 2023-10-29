package processor

import (
	"context"
	"log"

	"github.com/tencent-connect/botgo/dto"
)

func (p *Processor) setEmoji(ctx context.Context, channelID string, messageID string) {
	err := p.TencentAPI.OpenAPI.CreateMessageReaction(
		ctx, channelID, messageID, dto.Emoji{
			ID:   "307",
			Type: 1,
		},
	)
	if err != nil {
		log.Println(err)
	}
}

func (p *Processor) setPins(ctx context.Context, channelID, msgID string) {
	_, err := p.TencentAPI.OpenAPI.AddPins(ctx, channelID, msgID)
	if err != nil {
		log.Println(err)
	}
}

func (p *Processor) setAnnounces(ctx context.Context, data *dto.WSATMessageData) {
	if _, err := p.TencentAPI.OpenAPI.CreateChannelAnnounces(
		ctx, data.ChannelID,
		&dto.ChannelAnnouncesToCreate{MessageID: data.ID},
	); err != nil {
		log.Println(err)
	}
}

func (p *Processor) sendReply(ctx context.Context, channelID string, toCreate *dto.MessageToCreate) {
	if _, err := p.TencentAPI.OpenAPI.PostMessage(ctx, channelID, toCreate); err != nil {
		log.Println(err)
	}
}

func (p *Processor) sendReplyWithImg(ctx context.Context, channelID string, imgContents []byte, msgData map[string]string) {
	if _, err := p.TencentAPI.ExtendedAPI.PostMessageByFormData(ctx, channelID, imgContents, msgData); err != nil {
		log.Println(err)
	}
}
