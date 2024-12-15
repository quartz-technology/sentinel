package sentinel

import (
	"context"

	"github.com/quartz-technology/sentinel/pkg/services/decoder"
	"github.com/quartz-technology/sentinel/pkg/services/indexer"
	"github.com/quartz-technology/sentinel/pkg/services/listener"
	"github.com/quartz-technology/sentinel/pkg/services/notifier"
	"github.com/quartz-technology/sentinel/pkg/types"
	"golang.org/x/sync/errgroup"
)

type Service interface {
	Run(ctx context.Context) error
}

type service struct {
	eventsListenerService listener.EventsListener
	eventsDecoderService  decoder.Service
	indexerService        indexer.Service
	notifierService       notifier.Service
}

func NewService(
	eventsListenerService listener.EventsListener,
	eventsDecoderService decoder.Service,
	indexerService indexer.Service,
	notifierService notifier.Service,
) Service {
	return &service{
		eventsListenerService: eventsListenerService,
		eventsDecoderService:  eventsDecoderService,
		indexerService:        indexerService,
		notifierService:       notifierService,
	}
}

func (s *service) Run(ctx context.Context) error {
	capturedEventsLogs := make(chan *types.Log, 1)
	defer close(capturedEventsLogs)

	timelockedActions := make(chan types.TimelockedAction, 1)
	defer close(timelockedActions)

	eg, egCtx := errgroup.WithContext(ctx)
	defer eg.Wait()

	eg.Go(func() error {
		err := s.eventsListenerService.StartListeningForEventsLogs(egCtx, capturedEventsLogs)
		if err != nil {
			return err
		}

		return nil
	})

	eg.Go(func() error {
		err := s.eventsDecoderService.StartDecodingEventsLogs(egCtx, capturedEventsLogs, timelockedActions)
		if err != nil {
			return err
		}

		return nil
	})

	eg.Go(func() error {
		err := s.notifierService.StartNotifyingTimeLockedActions(egCtx, timelockedActions)
		if err != nil {
			return err
		}

		return nil
	})

	return nil
}