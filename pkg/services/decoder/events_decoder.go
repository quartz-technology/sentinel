package decoder

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/quartz-technology/sentinel/pkg/types"
	xcrypto "github.com/quartz-technology/sentinel/pkg/x/crypto"
)

// Service is used to decode the events that the listener service captured, which have been emitted by the MetaMorpho vaults.
//
// Those events are timelocked-actions, and the role of the decoder Service is to build a generic wrapper used later on
// by the indexer and notifier services.
type Service interface {
	// StartDecodingEventsLogs receives the captured events in a first channel, will decode them and send the result into a secondary channel.
	StartDecodingEventsLogs(ctx context.Context, eventsLogs <-chan *types.Log, timelockedActions chan<- types.TimelockedAction) error
}

// service is the implementation of the Service interface.
// It uses a mapping to associate each event signature hash to a specific timelocked-action type,
// which is used to wrap the decoded events.
type service struct {
	elClient                  *ethclient.Client
	metaMorphoVaultABI        abi.ABI
	sigHashToTimelockedAction map[string]types.TimelockedAction
}

func NewService(elClient *ethclient.Client, metaMorphoVaultABI abi.ABI) (Service, error) {
	eventSubmitTimelockSigHash, err := xcrypto.Keccak256Hash(metaMorphoVaultABI.Events["SubmitTimelock"].Sig)
	if err != nil {
		return nil, err
	}

	return &service{
		elClient:           elClient,
		metaMorphoVaultABI: metaMorphoVaultABI,
		sigHashToTimelockedAction: map[string]types.TimelockedAction{
			eventSubmitTimelockSigHash: &types.DecreaseTimelock{},
		},
	}, nil
}

func (s *service) StartDecodingEventsLogs(ctx context.Context, eventsLogs <-chan *types.Log, timelockedActions chan<- types.TimelockedAction) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case eventLog := <-eventsLogs:
			eventLogSigHash := eventLog.Value().Topics[0].String()

			timelockedAction, ok := s.sigHashToTimelockedAction[eventLogSigHash]

			if !ok {
				continue
			}

			timelockedAction, err := timelockedAction.DecodeFromEventLog(
				ctx,
				types.NewDecodingUtils(s.elClient, s.metaMorphoVaultABI),
				eventLog,
			)
			if err != nil {
				return err
			}

			timelockedActions <- timelockedAction
		}
	}
}
