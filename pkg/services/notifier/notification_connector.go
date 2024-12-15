package notifier

import (
	"context"

	"github.com/quartz-technology/sentinel/pkg/types"
	"github.com/sirupsen/logrus"
)

type NotificationConnector interface {
	SendNotification(ctx context.Context, action types.TimelockedAction) error
}

type terminalLogNotificationConnector struct {
}

func NewTerminalLogNotificationConnector() NotificationConnector {
	return &terminalLogNotificationConnector{}
}

func (c *terminalLogNotificationConnector) SendNotification(ctx context.Context, action types.TimelockedAction) error {
	logrus.
		WithFields(logrus.Fields{
			"vault":            action.VaultAddress(),
			"valid_at":         action.ValidAt(),
			"current_timelock": action.CurrentTimelock(),
			"description":      action.Debug(),
			"block_number":     action.BlockNumber(),
		}).
		Infoln("A new timelocked-action has been detected!")

	return nil
}
