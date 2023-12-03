package bot

import (
	"context"
	"fmt"
	"path"
	"runtime"
	"time"

	"github.com/Narushio/scarlet-bot/handler"
	"github.com/Narushio/scarlet-bot/pkg"
	"github.com/Narushio/scarlet-bot/pkg/token"
	"github.com/Narushio/scarlet-bot/pkg/websocket"
)

type ScarletBot struct {
	Version string
}

func NewScarletBot() *ScarletBot {
	return &ScarletBot{Version: "1.0.0"}
}

func (sb *ScarletBot) LinkStart() error {
	ctx := context.Background()
	botToken := token.New(token.TypeBot)
	if err := botToken.LoadFromConfig(getConfigPath("config.yaml")); err != nil {
		return err
	}

	api := pkg.NewOpenAPI(botToken).WithTimeout(3 * time.Second)

	wsInfo, err := api.WS(ctx, nil, "")
	if err != nil {
		return err
	}

	// websocket.RegisterResumeSignal(syscall.SIGUSR1)
	// 根据不同的回调，生成 intents
	intent := websocket.RegisterHandlers(
		// at 机器人事件，目前是在这个事件处理中有逻辑，会回消息，其他的回调处理都只把数据打印出来，不做任何处理
		handler.ATMessageEvent(),
		// 如果想要捕获到连接成功的事件，可以实现这个回调
		handler.Ready(),
		// 连接关闭回调
		handler.ErrorNotify(),
		// 频道事件
		handler.GuildEvent(),
		// 成员事件
		handler.MemberEvent(),
		// 子频道事件
		handler.ChannelEvent(),
		// 私信，目前只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		handler.DirectMessage(),
		// 频道消息，只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		handler.Message(ctx, api),
		// 互动事件
		handler.Interaction(),
		// 发帖事件
		handler.ThreadEvent(),
	)
	// 指定需要启动的分片数为 2 的话可以手动修改 wsInfo
	if err = pkg.NewSessionManager().Start(wsInfo, botToken, &intent); err != nil {
		return err
	}

	return nil
}

func getConfigPath(name string) string {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s/%s", path.Dir(filename), name)
	}
	return ""
}
