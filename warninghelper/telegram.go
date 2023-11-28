package warninghelper

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/Vealcoo/go-pkg/httphelper"
)

const (
	host    = "https://api.telegram.org/bot"
	timeOut = 5 * time.Second
)

type TelegramWarning struct {
	botToken string
	chatId   string
}

func NewTelegramWarning(token, chatId string) *TelegramWarning {
	t := new(TelegramWarning)
	t.botToken = token
	t.chatId = chatId

	return t
}

func (t *TelegramWarning) Ping() error {
	return t.sendMessage("for warning setup...")
}

func (t *TelegramWarning) Warning(msg interface{}) error {
	return t.sendMessage(msg)
}

type TelegramSendMessageRes struct {
	Ok bool `json:"ok"`
}

func (t *TelegramWarning) sendMessage(msg interface{}) error {
	apiHost := fmt.Sprintf("%s%s/%s", host, t.botToken, "sendMessage")

	req := url.Values{}
	req.Set("chat_id", t.chatId)
	str := fmt.Sprintf("%v", msg)
	req.Set("text", str)

	res := &TelegramSendMessageRes{}

	err := httphelper.SendPostWithUrl(apiHost, res, req, timeOut)
	if err != nil {
		return err
	}
	if !res.Ok {
		return errors.New("send telegram message failed")
	}

	return nil
}
