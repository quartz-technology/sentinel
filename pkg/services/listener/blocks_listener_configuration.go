package listener

import "time"

// BlocksListenerConfiguration is used to configure the Blocks Listener service.
type BlocksListenerConfiguration struct {
	// RefreshInterval is used by the polling implementation of the Service to delay two consecutive calls to eth_blockNumber.
	RefreshInterval time.Duration `mapstructure:"refresh-interval"`
}
