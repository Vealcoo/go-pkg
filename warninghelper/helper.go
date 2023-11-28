package warninghelper

import "errors"

type WarningClient int

const (
	// Auth with token and chatId
	Telegarm WarningClient = 1
)

func New(client WarningClient, auth ...string) (WarningService, error) {
	switch client {
	case Telegarm:
		tgClient := NewTelegramWarning(auth[0], auth[1])
		if err := tgClient.Ping(); err != nil {
			return nil, errors.New("warning setup failed")
		} else {
			return tgClient, nil
		}
	default:
		return nil, nil
	}
}

type WarningService interface {
	Warning(msg interface{}) error
}
