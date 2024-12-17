package logger

// Configuration is used to configure the logger client.
type Configuration struct {
	LogLevel string `mapstructure:"log-level" validate:"required"`
}
