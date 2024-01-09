package botgo

import (
	"github.com/Narushio/scarlet-bot/thirdparty/botgo/log"
	"github.com/Narushio/scarlet-bot/thirdparty/botgo/openapi"
	"github.com/Narushio/scarlet-bot/thirdparty/botgo/websocket"
)

// SetLogger 设置 logger，需要实现 internal 的 log.Logger 接口
func SetLogger(logger log.Logger) {
	log.DefaultLogger = logger
}

// SetSessionManager 注册自己实现的 session manager
func SetSessionManager(m SessionManager) {
	defaultSessionManager = m
}

// SetWebsocketClient 替换 websocket 实现
func SetWebsocketClient(c websocket.WebSocket) {
	websocket.Register(c)
}

// SetOpenAPIClient 注册 openapi 的不同实现，需要设置版本
func SetOpenAPIClient(v openapi.APIVersion, c openapi.OpenAPI) {
	openapi.Register(v, c)
}