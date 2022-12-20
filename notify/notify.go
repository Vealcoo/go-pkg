package notify

type Notify struct {
	tg *TelegramNotify
}

func New() *Notify {
	n := new(Notify)

	return n
}

func (n *Notify) SetTelegramNotify(token string) *Notify {
	t := NewTelegramNotify().SetToken(token)
	n.tg = t

	return n
}

func (n *Notify) NewTgClient() *TelegramNotify {
	if n.tg != nil {
		return n.tg
	}

	return nil
}
