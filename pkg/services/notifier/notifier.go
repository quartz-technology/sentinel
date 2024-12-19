package notifier

import (
	"context"

	"github.com/quartz-technology/sentinel/pkg/types"
	"github.com/sirupsen/logrus"
)

// Service is used to send notifications upon reception of a timelock action captured by Sentinel.
type Service interface {
	// StartNotifyingTimeLockedActions receives the decoded events and sends the notifications using connectors.
	StartNotifyingTimeLockedActions(ctx context.Context, timelockedActions <-chan types.TimelockedAction) error
}

// service is the implementation of the Service interface.
// It uses multiple connectors, each one implementing a specific communication channel.
type service struct {
	exitOnError bool
	connectors  []NotificationConnector
}

func NewService(config *Configuration, connectors []NotificationConnector) Service {
	exitOnError := false

	if config.ExitOnError != nil {
		exitOnError = *config.ExitOnError
	}

	return &service{
		exitOnError: exitOnError,
		connectors:  connectors,
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
					if s.exitOnError {
						return err
					}

					logrus.WithError(err).Errorln("failed to send notification using connector")
				}
			}
		}
	}
}
