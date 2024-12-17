package listener

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/quartz-technology/sentinel/pkg/bindings"
	"github.com/quartz-technology/sentinel/pkg/types"
	xcrypto "github.com/quartz-technology/sentinel/pkg/x/crypto"
	xiter "github.com/quartz-technology/sentinel/pkg/x/iter"
)

// EventsListener is used to capture the events emitted by the MetaMorpho Factory and Vaults at each block.
type EventsListener interface {
	// StartListeningForEventsLogs receives the captured blocks in a first channel, will capture the events logs emitted by the MetaMorpho contracts
	// and send the ones from the vaults into a secondary channel to be further processed by another service.
	StartListeningForEventsLogs(ctx context.Context, blocks <-chan uint64, eventsLogs chan<- *types.Log) error
}

// eventsListener is the implementation of the EventsListener interface.
// It will process the events emitted by the MetaMorpho Factory to also capture the events emitted by newly deployed Vaults.
type eventsListener struct {
	elClient                 *ethclient.Client
	metamorphoFactoryAddress common.Address
	metamorphoDeployedVaults []common.Address
	metaMorphoFactoryABI     abi.ABI
	metaMorphoVaultABI       abi.ABI
}

func NewEventsListener(
	config *EventsListenerConfiguration,
	elClient *ethclient.Client,
	metaMorphoFactoryABI,
	metaMorphoVaultABI abi.ABI,
) (EventsListener, error) {
	return &eventsListener{
		elClient:                 elClient,
		metamorphoFactoryAddress: common.HexToAddress(config.MetaMorphoFactoryAddress),
		metamorphoDeployedVaults: xiter.Map(config.MetaMorphoVaultAddresses, common.HexToAddress),
		metaMorphoFactoryABI:     metaMorphoFactoryABI,
		metaMorphoVaultABI:       metaMorphoVaultABI,
	}, nil
}

func (l *eventsListener) StartListeningForEventsLogs(ctx context.Context, blocks <-chan uint64, eventsLogs chan<- *types.Log) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case block := <-blocks:
			allTargetAddresses := []common.Address{l.metamorphoFactoryAddress}
			allTargetAddresses = append(allTargetAddresses, l.metamorphoDeployedVaults...)

			capturedEventsLogs, err := l.captureEventsLogsForBlock(ctx, block, allTargetAddresses)
			if err != nil {
				return err
			}

			for _, capturedEventLog := range capturedEventsLogs {
				eventsLogs <- capturedEventLog
			}
		}
	}
}

func (l *eventsListener) captureEventsLogsForBlock(ctx context.Context, block uint64, targetAddresses []common.Address) ([]*types.Log, error) {
	eventsLogs, err := l.fetchLogs(ctx, block, targetAddresses)
	if err != nil {
		return nil, err
	}

	newVaultsAddresses := make([]common.Address, 0)
	capturedEventsLogs := make([]*types.Log, 0)

	for _, eventLog := range eventsLogs {
		// NOTE: This event is emitted when a new vault is deployed.
		// We collect the new vault address to fetch events emitted from it right after finishing processing the already collected ones.
		if l.isMetaMorphoFactoryEventLog(eventLog) {
			eventName := "CreateMetaMorpho"
			isEventLogTarget, err := l.isEventLogTarget(eventLog, l.metaMorphoVaultABI.Events[eventName].Sig)
			if err != nil {
				return nil, err
			}

			if isEventLogTarget {
				event := &bindings.MetaMorphoFactoryCreateMetaMorpho{}

				err := l.metaMorphoFactoryABI.UnpackIntoInterface(event, eventName, eventLog.Data)
				if err != nil {
					return nil, err
				}

				newMetaMorphoVaultAddress := event.MetaMorpho

				newVaultsAddresses = append(newVaultsAddresses, newMetaMorphoVaultAddress)
				l.metamorphoDeployedVaults = append(l.metamorphoDeployedVaults, newMetaMorphoVaultAddress)
			}
		} else if l.isMetaMorphoVaultEventLog(eventLog) {
			capturedEventsLogs = append(
				capturedEventsLogs,
				types.NewLog(block, &eventLog),
			)
		}
	}

	// NOTE: Recursively fetch events from newly deployed vaults - if there are ones.
	if len(newVaultsAddresses) > 0 {
		capturedEventsLogsForNewVaults, err := l.captureEventsLogsForBlock(ctx, block, newVaultsAddresses)
		if err != nil {
			return nil, err
		}

		capturedEventsLogs = append(capturedEventsLogs, capturedEventsLogsForNewVaults...)
	}

	return capturedEventsLogs, nil
}

func (l *eventsListener) isMetaMorphoFactoryEventLog(eventLog ethtypes.Log) bool {
	return eventLog.Address == l.metamorphoFactoryAddress
}

func (l *eventsListener) isMetaMorphoVaultEventLog(eventLog ethtypes.Log) bool {
	for _, vaultAddress := range l.metamorphoDeployedVaults {
		if eventLog.Address == vaultAddress {
			return true
		}
	}

	return false
}

func (l *eventsListener) isEventLogTarget(eventLog ethtypes.Log, targetEventSig string) (bool, error) {
	eventSigHash, err := xcrypto.Keccak256Hash(targetEventSig)
	if err != nil {
		return false, err
	}

	return eventLog.Topics[0].String() == eventSigHash, nil
}

func (l *eventsListener) fetchLogs(ctx context.Context, block uint64, addresses []common.Address) ([]ethtypes.Log, error) {
	logs, err := l.elClient.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: big.NewInt(0).SetUint64(block),
		ToBlock:   big.NewInt(0).SetUint64(block),
		Addresses: addresses,
	})
	if err != nil {
		return nil, err
	}

	return logs, nil
}
