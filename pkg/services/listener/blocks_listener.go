package listener

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// BlocksListener is used to capture the new blocks numbers as the chain progresses.
type BlocksListener interface {
	// StartListeningForNewBlocks sends the new block numbers into a channel read by another service.
	StartListeningForNewBlocks(ctx context.Context, blocks chan<- uint64) error
}

// blocksListener is the implementation of the Service interface.
// It uses an HTTP client and the polling strategy to refresh its view of the latest block.
type blocksListener struct {
	refreshInterval time.Duration
	elClient        *ethclient.Client
}

func NewBlocksListener(config *BlocksListenerConfiguration, elClient *ethclient.Client) BlocksListener {
	return &blocksListener{
		refreshInterval: config.RefreshInterval,
		elClient:        elClient,
	}
}

func (l *blocksListener) StartListeningForNewBlocks(ctx context.Context, blocks chan<- uint64) error {
	oldBlock, err := l.elClient.BlockNumber(ctx)
	if err != nil {
		return err
	}

	blocks <- oldBlock

	t := time.NewTicker(l.refreshInterval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-t.C:
			newBlock, err := l.elClient.BlockNumber(ctx)
			if err != nil {
				return err
			}

			if newBlock > oldBlock {
				for i := oldBlock + 1; i <= newBlock; i++ {
					blocks <- i
				}

				oldBlock = newBlock
			}
		}
	}
}
