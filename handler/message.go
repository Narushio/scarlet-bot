package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Narushio/scarlet-bot/pkg/dto"
	"github.com/Narushio/scarlet-bot/pkg/event"
	"github.com/Narushio/scarlet-bot/pkg/openapi"
)

// Message 处理消息事件
func Message(ctx context.Context, api openapi.OpenAPI) event.MessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSMessageData) error {
		errChan := make(chan error, 1)

		go func() {
			time.Sleep(2 * time.Second)
			errChan <- errors.New("something went wrong")
		}()

		err := <-errChan
		if err != nil {
			fmt.Printf("Error received: %s\n", err)
			return err
		}

		return nil
	}
}
