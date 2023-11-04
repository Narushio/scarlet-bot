package qqx5

import (
	"context"

	"github.com/Narushio/scarlet-bot/plugin/qqx5/model"
)

type BoostMap interface {
	*model.DuoIdolBoostMap | *model.SoloIdolBoostMap
}

func (p *Plugin) sendBoostMap(ctx context.Context, subCmd string, text string, channelID string) error {

	if _, err := p.API.ExtendedAPI.PostMessageByFormData(ctx, channelID, nil, map[string]string{}); err != nil {
		return err
	}
	return nil
}
