package notifier

type Configuration struct {
	ExitOnError                  *bool                                      `mapstructure:"exit-on-error" validate:"required,boolean"`
	DiscordNotificationConnector *DiscordNotificationConnectorConfiguration `mapstructure:"discord-notification-connector"`
}

type DiscordNotificationConnectorConfiguration struct {
	WebHookURL string `mapstructure:"webhook-url" validate:"required"`
}
