package qqx5

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Narushio/scarlet-bot/api"
	"github.com/Narushio/scarlet-bot/browser"
	"github.com/Narushio/scarlet-bot/plugin/qqx5/model"
	"github.com/tencent-connect/botgo/dto"
)

const (
	BoostMapUrl = "http://localhost:8000/qqx5/duo-boost-map.html"
)

type Plugin struct {
	TencentAPI *api.TencentAPI
}

func New(api *api.TencentAPI) *Plugin {
	return &Plugin{TencentAPI: api}
}

func (p *Plugin) SendReplay(ctx context.Context, content string, data *dto.WSMessageData, atData *dto.WSATMessageData) error {
	c := strings.Split(content, " ")
	cmdPrefix := c[0]
	content = c[1]
	switch cmdPrefix {
	case "双排星动爆点":
		if err := p.sendBoostMap(ctx, content, data, atData, model.Idol, model.Duo); err != nil {
			return err
		}
	}
	return nil
}

func (p *Plugin) sendBoostMap(ctx context.Context, content string, data *dto.WSMessageData, atData *dto.WSATMessageData, pt model.Pattern, m model.Mode) error {
	var channelID string
	var msgID string
	var authorID string

	if data != nil {
		channelID = data.ChannelID
		msgID = data.ID
		authorID = data.Author.ID
	}
	if atData != nil {
		channelID = atData.ChannelID
		msgID = atData.ID
		authorID = atData.Author.ID
	}

	page, err := browser.Chromiun.NewPage()
	if err != nil {
		return err
	}
	if _, err = page.Goto(BoostMapUrl); err != nil {
		return err
	}

	var domData []byte
	if pt == model.Idol && m == model.Duo {
		var dbm *model.DuoBoostMap
		for _, d := range model.DuoBoostMaps {
			if strings.Contains(strings.ToLower(d.Title), content) {
				dbm = d
			}
		}
		domData, err = json.Marshal(dbm)
		if err != nil {
			return err
		}
	}

	_, err = page.Evaluate(fmt.Sprintf("initDom(%s)", string(domData)))
	if err != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)
	imgContents, err := page.Locator("div.main").Screenshot()
	if err != nil {
		return err
	}

	msgData := map[string]string{
		"content": fmt.Sprintf("<@%s>", authorID),
		"msg_id":  msgID,
	}

	if _, err := p.TencentAPI.ExtendedAPI.PostMessageByFormData(ctx, channelID, imgContents, msgData); err != nil {
		return err
	}

	return nil
}
