package listener

import "time"

type BlocksListenerConfiguration struct {
	RefreshInterval time.Duration `mapstructure:"refresh-interval"`
}
