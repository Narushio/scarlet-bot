package action

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Narushio/scarlet-bot/browser"
	"github.com/Narushio/scarlet-bot/plugin/qqx5/model"
)

type BoostMap interface {
	*model.DuoIdolBoostMap | *model.SoloIdolBoostMap
	GetTitle() string
	GetTemplateUrl() string
}

//func SendDuoIdolBoostMap() *bot.CmdAction {
//cmds := []string{"双排星动爆点", "星动双排爆点", "星动双排", "双排星动"}
//action := func(ctx context.Context, api *openapi.OpenAPI, messageData any) error {
//	imgData, err := buildBoostMapImg[*model.DuoIdolBoostMap]("双排", model.DuoIdolBoostMapTemplateUrl, model.DuoIdolBoostMapList)
//
//	if err := api.ReplyWithImg(ctx, messageData, nil, nil); err != nil {
//		return err
//	}
//	return nil
//}
//return bot.NewCmdAction("sendDuoIdolBoostMap", cmds, action)
//}

func buildBoostMapImg[T BoostMap](text string, templateUrl string, mapList []T) ([]byte, error) {
	page, err := browser.Chromium.NewPage()
	if err != nil {
		return nil, err
	}

	var domData []byte
	for _, m := range mapList {
		if strings.Contains(strings.ToLower(m.GetTitle()), text) {
			domData, err = json.Marshal(m)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	if _, err = page.Goto(templateUrl); err != nil {
		return nil, err
	}

	if _, err = page.Evaluate(fmt.Sprintf("initDom(%s)", string(domData))); err != nil {
		return nil, err
	}

	time.Sleep(500 * time.Millisecond)
	imgData, err := page.Locator("div.main").Screenshot()
	if err != nil {
		return nil, err
	}

	return imgData, nil
}
