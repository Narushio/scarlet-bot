package bot

import (
	"context"

	"github.com/Narushio/scarlet-bot/pkg/dto"
	"github.com/Narushio/scarlet-bot/pkg/openapi"
)

type Action func(ctx context.Context, api *openapi.OpenAPI, data *dto.WSPayload) error

type CmdAction struct {
	Name       string
	InvokeCmds []string
	Action     Action
}

func NewCmdAction(name string, cmds []string, action Action) *CmdAction {
	return &CmdAction{
		Name:       name,
		InvokeCmds: cmds,
		Action:     action,
	}
}
