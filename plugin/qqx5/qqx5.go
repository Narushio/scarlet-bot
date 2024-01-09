package qqx5

import (
	"context"
	"strings"

	"github.com/Narushio/scarlet-bot/thirdparty/botgo/dto"
	"github.com/Narushio/scarlet-bot/thirdparty/botgo/openapi"
)

const (
	ResourcePath = "resource/qqx5"
	InvokeCmd    = "/x5"
)

func init() {

}

type Plugin struct {
	Version   string
	Name      string
	InvokeCmd string
}

func New() *Plugin {
	return &Plugin{
		Version:   "1.0.0",
		Name:      "QQ炫舞手游",
		InvokeCmd: "qqx5",
	}
}

func (p *Plugin) Reply(ctx context.Context, api *openapi.OpenAPI, event *dto.WSPayload, errChan chan error) {
	content := event.MessageData.Content
	if !strings.HasPrefix(content, p.InvokeCmd) {
		return
	}

	subCmd, text := parseCmdContent(content)
	for _, action := range p.Actions {
		if action.Name == subCmd {
			action.Run(ctx, api, event, errChan)
			return
		}
	}
}

func parseCmdContent(content string) (subCmd string, text string) {
	c := strings.Split(content, " ")
	subCmd = c[0]
	text = c[1]
	return
}
