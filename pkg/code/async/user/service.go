package async_user

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/code-payments/code-server/pkg/code/async"
	code_data "github.com/code-payments/code-server/pkg/code/data"
	push_lib "github.com/code-payments/code-server/pkg/push"
	"github.com/code-payments/code-server/pkg/sync"
	"github.com/code-payments/code-server/pkg/twitter"
)

type service struct {
	log           *logrus.Entry
	data          code_data.Provider
	pusher        push_lib.Provider
	twitterClient *twitter.Client
	userLocks     *sync.StripedLock
}

func New(data code_data.Provider, pusher push_lib.Provider, twitterClient *twitter.Client) async.Service {
	return &service{
		log:           logrus.StandardLogger().WithField("service", "user"),
		data:          data,
		pusher:        pusher,
		twitterClient: twitterClient,
		userLocks:     sync.NewStripedLock(1024),
	}
}

// todo: split out interval for each worker
func (p *service) Start(ctx context.Context, interval time.Duration) error {
	go func() {
		err := p.twitterRegistrationWorker(ctx, interval)
		if err != nil && err != context.Canceled {
			p.log.WithError(err).Warn("twitter registration processing loop terminated unexpectedly")
		}
	}()

	go func() {
		err := p.twitterUserInfoUpdateWorker(ctx, interval)
		if err != nil && err != context.Canceled {
			p.log.WithError(err).Warn("twitter user info processing loop terminated unexpectedly")
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}
