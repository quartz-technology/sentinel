package elclient

// Configuration is used to configure the Ethereum Execution Layer RPC client.
type Configuration struct {
	// RPCURL is the Ethereum node endpoint.
	RPCURL string `mapstructure:"rpc-url" validate:"required,uri"`
}
