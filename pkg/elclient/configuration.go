package elclient

type Configuration struct {
	RPCURL string `mapstructure:"rpc-url" validate:"required,uri"`
}
