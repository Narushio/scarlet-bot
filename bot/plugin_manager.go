package bot

import (
	"context"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

type Plugin interface {
	VerifyCmd(cmd message.CMD) bool
	HandleMessageData(ctx context.Context, data *dto.WSMessageData) error
	HandleATMessageData(ctx context.Context, data *dto.WSATMessageData) error
	HandleDirectMessageData(ctx context.Context, data *dto.WSDirectMessageData) error
}

type PluginManager struct {
	PluginList []Plugin
}

func (pm *PluginManager) RegisterPlugin(plugin Plugin) {
	pm.PluginList = append(pm.PluginList, plugin)
}
