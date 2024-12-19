package notifier

import (
	"context"
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"github.com/quartz-technology/sentinel/pkg/types"
	"github.com/sirupsen/logrus"
)

// A NotificationConnector is used to connect to a service able to receive the notification (Discord, Telegram, Slack, ...).
type NotificationConnector interface {
	// SendNotification sends the notification alerting the detection of a timelocked-action to the user.
	SendNotification(ctx context.Context, action types.TimelockedAction) error
}

type terminalLogNotificationConnector struct{}

func NewTerminalLogNotificationConnector() NotificationConnector {
	return &terminalLogNotificationConnector{}
}

func (c *terminalLogNotificationConnector) SendNotification(ctx context.Context, action types.TimelockedAction) error {
	logrus.
		WithFields(logrus.Fields{
			"kind":             action.Kind(),
			"vault":            action.VaultAddress(),
			"valid_at":         action.ValidAt(),
			"current_timelock": action.CurrentTimelock(),
			"description":      action.Debug(),
			"chain":            action.VaultNetwork().Short,
			"block_number":     action.BlockNumber(),
		}).
		Infoln("A new timelocked-action has been detected!")

	return nil
}

type discordNotificationConnector struct {
	client webhook.Client
}

func NewDiscordNotificationConnector(config *DiscordNotificationConnectorConfiguration) (NotificationConnector, error) {
	client, err := webhook.NewWithURL(config.WebHookURL)
	if err != nil {
		return nil, err
	}

	return &discordNotificationConnector{
		client: client,
	}, nil
}

func (c *discordNotificationConnector) SendNotification(ctx context.Context, action types.TimelockedAction) error {
	_, err := c.client.CreateMessage(
		discord.
			NewWebhookMessageCreateBuilder().
			SetContent(fmt.Sprintf(
				`# :warning:  New Timelocked-Action Detected :warning:

:gear: **Action Kind**: [%s]

:bank: **Vault**: %s

:alarm_clock: **Valid At**: %s

:timer: **Current Timelock**: %s

:question:**Description**: %s

:chains: **Chain**: %s

:bricks: **Block Number**: %d

:link: **Link to vault**: https://app.morpho.org/vault?vault=%s&network=%s`,
				action.Kind(),
				action.VaultAddress(),
				action.ValidAt(),
				action.CurrentTimelock(),
				action.Debug(),
				action.VaultNetwork().Long,
				action.BlockNumber(),
				action.VaultAddress(), action.VaultNetwork().Short,
			)).
			Build(),
	)
	if err != nil {
		return err
	}

	return nil
}
