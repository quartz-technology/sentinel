package types

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DecodingUtils struct {
	elClient           *ethclient.Client
	metaMorphoVaultABI abi.ABI
}

func NewDecodingUtils(elClient *ethclient.Client, metaMorphoVaultABI abi.ABI) *DecodingUtils {
	return &DecodingUtils{
		elClient:           elClient,
		metaMorphoVaultABI: metaMorphoVaultABI,
	}
}

func (u *DecodingUtils) ELClient() *ethclient.Client {
	return u.elClient
}

func (u *DecodingUtils) MetaMorphoVaultABI() abi.ABI {
	return u.metaMorphoVaultABI
}
