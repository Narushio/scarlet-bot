package qqx5

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/tencent-connect/botgo/dto"

	"github.com/Narushio/scarlet-bot/api"
	"github.com/Narushio/scarlet-bot/browser"
	"github.com/Narushio/scarlet-bot/plugin/qqx5/loader"
	"github.com/Narushio/scarlet-bot/plugin/qqx5/model"
)

const (
	DuoIdolBoostMapUrl  = "http://localhost:8000/qqx5/duo-idol-boost-map.html"
	SoloIdolBoostMapUrl = "http://localhost:8000/qqx5/solo-idol-boost-map.html"
)

func init() {
	loader.BoostMapExcel()
}

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
	case "单排星动爆点":
		if err := p.sendBoostMap(ctx, content, data, atData, model.Idol, model.Solo); err != nil {
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

	domData, err := fetchDomData(pt, m, content, page)
	if err != nil {
		return err
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

func fetchDomData(pt model.Pattern, m model.Mode, content string, page playwright.Page) ([]byte, error) {
	var url string
	var mapList []interface{}
	var err error
	switch {
	case pt == model.Idol && m == model.Duo:
		url = DuoIdolBoostMapUrl
		mapList = interfaceSlice(model.DuoIdolBoostMapList)
	case pt == model.Idol && m == model.Solo:
		url = SoloIdolBoostMapUrl
		mapList = interfaceSlice(model.SoloIdolBoostMapList)
	default:
		return nil, errors.New("invalid pattern or mode")
	}

	var mapData []byte
	for _, d := range mapList {
		title := reflect.ValueOf(d).Elem().FieldByName("Title").String()
		if strings.Contains(strings.ToLower(title), content) {
			mapData, err = json.Marshal(d)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	if _, err = page.Goto(url); err != nil {
		return nil, err
	}

	return mapData, nil
}

func interfaceSlice(slice interface{}) []interface{} {
	values := reflect.ValueOf(slice)
	if values.Kind() != reflect.Slice {
		return nil
	}

	result := make([]interface{}, values.Len())
	for i := 0; i < values.Len(); i++ {
		result[i] = values.Index(i).Interface()
	}
	return result
}
