package sentinel

import (
	"github.com/quartz-technology/sentinel/pkg/elclient"
	"github.com/quartz-technology/sentinel/pkg/services/listener"
	"github.com/quartz-technology/sentinel/pkg/services/logger"
)

// Configuration is used to configure the Sentinel Service.
type Configuration struct {
	Logger         *logger.Configuration                 `mapstructure:"logger" validate:"required"`
	ELClient       *elclient.Configuration               `mapstructure:"el-client" validate:"required"`
	BlocksListener *listener.BlocksListenerConfiguration `mapstructure:"blocks-listener" validate:"required"`
	EventsListener *listener.EventsListenerConfiguration `mapstructure:"events-listener" validate:"required"`
}
