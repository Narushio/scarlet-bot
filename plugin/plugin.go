package plugin

import (
	"context"

	"github.com/Narushio/scarlet-bot/pkg/dto"
	"github.com/Narushio/scarlet-bot/pkg/openapi"
)

type Action func(ctx context.Context, api *openapi.OpenAPI, event *dto.WSPayload, errChan chan error)

type Plugin interface {
	baseInfo
	Reply(ctx context.Context, api *openapi.OpenAPI, event *dto.WSPayload, errChan chan error)
}

type baseInfo struct {
	Version   string
	Name      string
	InvokeCmd string
}
