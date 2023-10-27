package api

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
)

type ExtendedAPI interface {
	SendPicToChannelMsg(ctx context.Context, channelID string, filename string, data map[string]string) ([]byte, error)
}

type extendedAPI struct {
	token       *token.Token
	timeout     time.Duration
	restyClient *resty.Client
}

type TencentAPI struct {
	OpenAPI     openapi.OpenAPI
	ExtendedAPI ExtendedAPI
}

func NewTencentAPI(botToken *token.Token) *TencentAPI {
	oAPI := botgo.NewSandboxOpenAPI(botToken)
	eAPI := &extendedAPI{
		token: botToken,
	}
	eAPI.setupClient()
	return &TencentAPI{
		OpenAPI:     oAPI,
		ExtendedAPI: eAPI,
	}
}

func (t *TencentAPI) WithTimeout(duration time.Duration) *TencentAPI {
	t.OpenAPI.WithTimeout(duration)
	t.ExtendedAPI.(*extendedAPI).restyClient.SetTimeout(duration)
	return t
}

func (e *extendedAPI) setupClient() {
	e.restyClient = resty.New().
		SetAuthScheme(string(e.token.Type)).
		SetAuthToken(e.token.GetString()).
		SetContentLength(true)
}

func (e *extendedAPI) SendPicToChannelMsg(ctx context.Context, channelID string, filename string, data map[string]string) ([]byte, error) {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	resp, err := e.restyClient.R().
		SetContext(ctx).
		SetFormData(data).
		SetResult(dto.Message{}).
		SetPathParam("channel_id", channelID).
		SetFileReader("file_image", filename, bytes.NewReader(fileContents)).
		Post(fmt.Sprintf("%s://%s%s", "https", "sandbox.api.sgroup.qq.com", "/channels/{channel_id}/messages"))
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
