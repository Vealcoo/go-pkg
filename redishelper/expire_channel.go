package redishelper

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

type ExpiredChannel struct {
	client *redis.Client
	action map[string]func(c *ExpiredChannelContext) // key:action, value:function

	ctx *ExpiredChannelContext
}

func SetExpiredChannel(client *redis.Client) *ExpiredChannel {
	return &ExpiredChannel{
		client: client,
	}
}

type ExpiredChannelContext struct {
	// TODO: 放連線資訊之類的

	info string
}

func (e *ExpiredChannelContext) setInfo(info string) {
	e.info = info
}

func (e *ExpiredChannelContext) GetInfo() string {
	return e.info
}

func (e *ExpiredChannel) SubExpiredEvent() {
	ctx := context.Background()
	expirer := e.client.Subscribe(ctx, "__keyevent@0__:expired")
	defer e.client.Close()

	for {
		msg, err := expirer.ReceiveMessage(ctx)
		if err != nil {
			log.Error().Msg(err.Error())
		}

		index := strings.Index(msg.Payload, "/")
		if index == -1 {
			continue
		}

		action := msg.Payload[:index]
		info := msg.Payload[index+1:]
		e.ctx.setInfo(info)

		if _, ok := e.action[action]; ok {
			e.action["action"](e.ctx)
		}
	}
}

func (e *ExpiredChannel) AddAction(action string, f func(c *ExpiredChannelContext)) {
	e.action[action] = f
}

func (e *ExpiredChannel) CancelAction(actionName string) error {
	if _, ok := e.action[actionName]; !ok {
		return errors.New("no such action")
	}
	delete(e.action, actionName)

	return nil
}

// 先加 action 在加 event
func (e *ExpiredChannel) SetEvent(actionName, info string, eventTime int64) error {
	if _, ok := e.action[actionName]; !ok {
		return errors.New("no such action")
	}

	// FIXME: 驗證一下 actionName 裡面不要有 "/" ~~~~~
	key := fmt.Sprintf("%s/%s", actionName, info)
	return e.client.Set(context.Background(), key, "for expired event", time.Duration(eventTime)*time.Second).Err()
}

