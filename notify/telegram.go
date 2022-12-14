package notify

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	httpHelper "github.com/Vealcoo/go-pkg/http_helper"
)

const (
	host    = "https://api.telegram.org/bot"
	timeOut = 5 * time.Second
)

type TelegramNotify struct {
	botToken string
}

func NewTelegramNotify() *TelegramNotify {
	t := new(TelegramNotify)

	return t
}

func (t *TelegramNotify) SetToken(token string) *TelegramNotify {
	t.botToken = token

	return t
}

type TelegramNotifyRes struct {
	Ok bool `json:"ok"`
}

func (t *TelegramNotify) Notify(msg interface{}, chatId int64) (*TelegramNotifyRes, error) {
	apiHost := fmt.Sprintf("%s%s/%s", host, t, "sendMessage")

	req := url.Values{}
	req.Set("chat_id", strconv.FormatInt(chatId, 10))
	str := fmt.Sprintf("%v", msg)
	req.Set("text", str)

	res := &TelegramNotifyRes{}

	err := httpHelper.SendPostWithUrl(apiHost, res, req, timeOut)
	if err != nil {
		return nil, err
	}

	return res, nil
}