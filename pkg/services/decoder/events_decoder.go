package decoder

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/quartz-technology/sentinel/pkg/types"
	xcrypto "github.com/quartz-technology/sentinel/pkg/x/crypto"
)

type Service interface {
	StartDecodingEventsLogs(ctx context.Context, eventsLogs <-chan *types.Log, timelockedActions chan<- types.TimelockedAction) error
}

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
