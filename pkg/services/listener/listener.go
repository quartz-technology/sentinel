package listener

import (
	"context"

	"github.com/quartz-technology/sentinel/pkg/types"
)

type Service interface {
	StartListeningForNewTimelockedActions(ctx context.Context, actions chan<- types.TimelockedAction) error
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) StartListeningForNewTimelockedActions(ctx context.Context, actions chan<- types.TimelockedAction) error {
	// TODO: Listen for new blocks.

	// TODO: Listen for new deployed vaults.

	// TODO: Listen for emitted events for all deployed vaults in the current block.

	return nil
}

//
// func NewService(conf *Configuration) Service {
// 	return &service{
// 		conf: conf,
// 	}
// }
//
// func (s *service) Listen(ctx context.Context, actions chan<- types.TimelockedAction) error {
// 	contractAbi, err := abi.JSON(strings.NewReader(string(bindings.EventLibABI)))
// 	if err != nil {
// 		return err
// 	}
//
// 	eg, egCtx := errgroup.WithContext(ctx)
// 	defer eg.Wait()
//
// 	for _, listeningConf := range s.conf.VaultAddressesByChain {
// 		eg.Go(func() error {
// 			client, err := ethclient.Dial(listeningConf.WebSocketEndpoint)
// 			if err != nil {
// 				return err
// 			}
//
// 			ch := make(chan ethtypes.Log)
// 			sub, err := client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{
// 				Addresses: listeningConf.VaultAddresses,
// 			}, ch)
// 			if err != nil {
// 				return err
// 			}
//
// 			for {
// 				select {
// 				case <-egCtx.Done():
// 					sub.Unsubscribe()
//
// 					return egCtx.Err()
// 				case eventLog := <-ch:
// 					e := &bindings.EventLibSubmitTimelock{}
//
// 					err := contractAbi.UnpackIntoInterface(e, "SubmitTimelock", eventLog.Data)
// 					if err != nil {
// 						return err
// 					}
//
// 					log.Printf("Event [SubmitTimelock]: %v", e)
// 				}
// 			}
// 		})
// 	}
//
// 	return nil
// }

// func (s *service) ListenWithCallback(ctx context.Context, callback func(action types.TimelockedAction) error) error {
// 	return nil
// }
