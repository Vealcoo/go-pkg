package fcmhelper

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/pkg/errors"
	zerolog "github.com/rs/zerolog"
)

type fcmSender struct {
	client  *messaging.Client
	log     zerolog.Logger
	workers []*fcmWorker
	c       chan *messaging.MulticastMessage
}

func CreateFCMSender(ctx context.Context, workerNum int, l zerolog.Logger) *fcmSender {
	firebaseApp, err := firebase.NewApp(ctx, nil)
	if err != nil {
		l.Panic().Err(err).Send()
	}

	messageClient, err := firebaseApp.Messaging(ctx)
	if err != nil {
		l.Panic().Err(err).Send()
	}

	if workerNum == 0 {
		l.Panic().Err(errors.New("workerNum is 0")).Send()
	}

	s := &fcmSender{
		client:  messageClient,
		log:     l,
		workers: make([]*fcmWorker, workerNum),
		c:       make(chan *messaging.MulticastMessage, 1024),
	}

	for i := 0; i < workerNum; i++ {
		s.workers[i] = &fcmWorker{s}
		go s.workers[i].start()
	}

	return s
}

func (s *fcmSender) Send(d *messaging.MulticastMessage) {
	s.c <- d
}

type fcmWorker struct {
	sender *fcmSender
}

func (w *fcmWorker) start() {
	for m := range w.sender.c {
		res, err := w.sender.client.SendMulticast(context.Background(), m)
		if err != nil {
			w.sender.log.Error().Stack().Err(errors.WithStack(err)).Send()
			continue
		}

		w.sender.log.Info().Int("success", res.SuccessCount).Int("fail", res.FailureCount).Interface("responses", res.Responses).Strs("tokens", m.Tokens).Msg("FCM send")
	}
}
