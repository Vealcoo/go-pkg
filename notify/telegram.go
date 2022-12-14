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

type TelegramNotifyRes struct {
	Ok bool `json:"ok"`
}

func Notify(botKey, msg string, chatId int64) (*TelegramNotifyRes, error) {
	apiHost := fmt.Sprintf("%s%s/%s", host, botKey, "sendMessage")

	req := url.Values{}
	req.Set("chat_id", strconv.FormatInt(chatId, 10))
	req.Set("text", msg)

	res := &TelegramNotifyRes{}

	err := httpHelper.SendPostWithUrl(apiHost, res, req, timeOut)
	if err != nil {
		return nil, err
	}

	return res, nil
}
