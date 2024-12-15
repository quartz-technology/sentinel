package types

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/quartz-technology/sentinel/pkg/bindings"
)

type DecreaseTimelock struct {
	*BaseTimelockedAction

	newTimelock *uint256.Int
}

func NewDecreaseTimelockAction(base *BaseTimelockedAction, newTimelock *uint256.Int) TimelockedAction {
	return &DecreaseTimelock{
		BaseTimelockedAction: base,
		newTimelock:          newTimelock,
	}
}

func (a *DecreaseTimelock) DecodeFromEventLog(ctx context.Context, utils *DecodingUtils, eventLog *Log) (TimelockedAction, error) {
	event := &bindings.MetaMorphoVaultSubmitTimelock{}

	err := utils.MetaMorphoVaultABI().UnpackIntoInterface(event, "SubmitTimelock", eventLog.value.Data)
	if err != nil {
		return nil, err
	}

	newTimelock, overflow := uint256.FromBig(event.NewTimelock)
	if overflow {
		return nil, fmt.Errorf("overflow when converting new timelock to uint256")
	}

	metaMorphoVaultInstance, err := bindings.NewMetaMorphoVault(eventLog.value.Address, utils.elClient)
	if err != nil {
		return nil, err
	}

	res, err := metaMorphoVaultInstance.PendingTimelock(&bind.CallOpts{
		BlockNumber: big.NewInt(0).SetUint64(eventLog.blockNumber),
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	currentTimelock, err := metaMorphoVaultInstance.Timelock(&bind.CallOpts{
		BlockNumber: big.NewInt(0).SetUint64(eventLog.blockNumber),
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return NewDecreaseTimelockAction(
		NewBaseTimelockedAction(eventLog.value.Address, currentTimelock, res.ValidAt, eventLog.blockNumber),
		newTimelock,
	), nil
}

func (a *DecreaseTimelock) VaultAddress() common.Address {
	return a.vaultAddress
}

func (a *DecreaseTimelock) CurrentTimelock() time.Duration {
	return time.Duration(a.currentTimelock.Uint64()) * time.Second
}

func (a *DecreaseTimelock) ValidAt() time.Time {
	return time.Unix(int64(a.validAt), 0)
}

func (a *DecreaseTimelock) BlockNumber() uint64 {
	return a.blockNumber
}

func (a *DecreaseTimelock) Debug() string {
	newTimelockDuration := time.Duration(a.newTimelock.Uint64()) * time.Second

	return fmt.Sprintf("Timelock decreased to %s", newTimelockDuration)
}
