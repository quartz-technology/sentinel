package notifier

import (
	"context"

	"github.com/quartz-technology/sentinel/pkg/types"
)

type Service interface {
	StartNotifyingTimeLockedActions(ctx context.Context, timelockedActions <-chan types.TimelockedAction) error
}

type service struct {
	connectors []NotificationConnector
}

func NewService(connectors []NotificationConnector) Service {
	return &service{
		connectors: connectors,
	}
}

func (s *service) StartNotifyingTimeLockedActions(ctx context.Context, timelockedActions <-chan types.TimelockedAction) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case timelockedAction := <-timelockedActions:
			for _, connector := range s.connectors {
				if err := connector.SendNotification(ctx, timelockedAction); err != nil {
					return err
				}
			}
		}
	}
}
