package listener

// EventsListenerConfiguration is used to configure the Events Listener service.
type EventsListenerConfiguration struct {
	// MetaMorphoFactoryAddress is the address of the MetaMorpho Factory contract.
	MetaMorphoFactoryAddress string `mapstructure:"metamorpho-factory-address" validate:"required,eth_addr"`

	// MetaMorphoVaultAddresses are the addresses of the MetaMorpho Vaults contract already deployed.
	MetaMorphoVaultAddresses []string `mapstructure:"metamorpho-vault-addresses" validate:"required,dive,eth_addr"`
}
