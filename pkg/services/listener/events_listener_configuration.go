package listener

type EventsListenerConfiguration struct {
	MetaMorphoFactoryAddress string   `mapstructure:"metamorpho-factory-address" validate:"required,eth_addr"`
	MetaMorphoVaultAddresses []string `mapstructure:"metamorpho-vault-addresses" validate:"required,dive,eth_addr"`
}
