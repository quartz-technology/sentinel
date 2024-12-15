package types

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type TimelockedAction interface {
	DecodeFromEventLog(ctx context.Context, utils *DecodingUtils, eventLog *Log) (TimelockedAction, error)
	VaultAddress() common.Address
	CurrentTimelock() time.Duration
	ValidAt() time.Time
	BlockNumber() uint64
	Debug() string
}

type BaseTimelockedAction struct {
	vaultAddress    common.Address
	currentTimelock *big.Int
	validAt         uint64
	blockNumber     uint64
}

func NewBaseTimelockedAction(vaultAddress common.Address, currentTimelock *big.Int, validAt uint64, blockNumber uint64) *BaseTimelockedAction {
	return &BaseTimelockedAction{
		vaultAddress:    vaultAddress,
		currentTimelock: currentTimelock,
		validAt:         validAt,
		blockNumber:     blockNumber,
	}
}
