package logger

type Configuration struct {
	LogLevel string `mapstructure:"log-level" validate:"required"`
}
