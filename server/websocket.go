package main

import (
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	// WebSocket 服务器地址
	serverAddr := "ws://127.0.0.1:5701"

	// 使用默认的 WebSocket 握手选项连接到服务器
	c, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatalf("Failed to connect to the WebSocket server: %v", err)
	}
	defer c.Close()

	// 以下是一个简单的例子，从服务器读取并打印消息
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)
	}
}
