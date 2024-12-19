package types

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DecodingUtils struct {
	elClient           *ethclient.Client
	metaMorphoVaultABI abi.ABI
	networks           map[uint64]NetworkName
}

func NewDecodingUtils(elClient *ethclient.Client, metaMorphoVaultABI abi.ABI) *DecodingUtils {
	return &DecodingUtils{
		elClient:           elClient,
		metaMorphoVaultABI: metaMorphoVaultABI,
		networks: map[uint64]NetworkName{
			1: {
				Long:  "Ethereum Mainnet",
				Short: "mainnet",
			},
			8453: {
				Long:  "Base Mainnet",
				Short: "base",
			},
		},
	}
}

func (u *DecodingUtils) ELClient() *ethclient.Client {
	return u.elClient
}

func (u *DecodingUtils) MetaMorphoVaultABI() abi.ABI {
	return u.metaMorphoVaultABI
}

func (u *DecodingUtils) Network(id uint64) NetworkName {
	return u.networks[id]
}
