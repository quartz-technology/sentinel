package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-playground/validator/v10"
	"github.com/quartz-technology/sentinel/pkg/bindings"
	"github.com/quartz-technology/sentinel/pkg/elclient"
	"github.com/quartz-technology/sentinel/pkg/services/decoder"
	"github.com/quartz-technology/sentinel/pkg/services/indexer"
	"github.com/quartz-technology/sentinel/pkg/services/listener"
	"github.com/quartz-technology/sentinel/pkg/services/logger"
	"github.com/quartz-technology/sentinel/pkg/services/notifier"
	"github.com/quartz-technology/sentinel/pkg/services/sentinel"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const sentinelAsciiArt = `
     _____            __  _            __
    / ___/___  ____  / /_(_)___  ___  / /
    \__ \/ _ \/ __ \/ __/ / __ \/ _ \/ / 
   ___/ /  __/ / / / /_/ / / / /  __/ /  
  /____/\___/_/ /_/\__/_/_/ /_/\___/_/   

                                         `

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sentinel",
	Short: "Indexing, monitoring and alerting solution for Morpho vaults time-locked operations.",
	Long: fmt.Sprintf(`%s
Sentinel is a tool that helps the users of MetaMorpho vaults to get realtime alerts whenever the vault curator performs a timelocked operation, 
i.e. an action that has a potential impact on the user deposit.

There are different types of timelocked actions:
 - Decrease the timelock duration.
 - Set a vault Guardian (if not already set).
 - Increase the supply cap of any market in a vault.
 - Submit the forced removal of a market from a vault.

Using various connectors, the users can receive alerts on platforms such as Discord, Slack, Telegram and so on.
You can also configure a custom connector that fits your own needs!
`, sentinelAsciiArt),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		defer stop()

		sentinelConfig := new(sentinel.Configuration)
		if err := viper.Unmarshal(sentinelConfig); err != nil {
			return err
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(sentinelConfig)

		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			if validationErrors != nil {
				return validationErrors
			}
		}

		initLogger(sentinelConfig.Logger)

		elClient, err := initELClient(ctx, sentinelConfig.ELClient)
		if err != nil {
			return err
		}

		blocksListenerService := initBlocksListenerService(sentinelConfig.BlocksListener, elClient)

		metaMorphoFactoryABI, metaMorphoVaultABI, err := initABIs()
		if err != nil {
			return err
		}

		eventsListenerService, err := initEventsListenerService(
			sentinelConfig.EventsListener,
			elClient,
			metaMorphoFactoryABI,
			metaMorphoVaultABI,
		)
		if err != nil {
			return err
		}

		eventsDecoderService, err := initDecoderService(
			elClient,
			metaMorphoVaultABI,
		)
		if err != nil {
			return err
		}

		indexerService := initIndexerService()

		notifierService := initNotifierService()

		sentinelService := sentinel.NewService(
			blocksListenerService,
			eventsListenerService,
			eventsDecoderService,
			indexerService,
			notifierService,
		)

		logrus.Debug("All services initialized, starting sentinel..")

		err = sentinelService.Run(ctx)
		if err != nil {
			return err
		}

		logrus.Debug("Sentinel interrupted, shutting down..")

		return nil
	},
}

func initLogger(loggerConfig *logger.Configuration) error {
	logLevel := loggerConfig.LogLevel

	var logLevelValue logrus.Level

	switch logLevel {
	case "panic":
		logLevelValue = logrus.PanicLevel
	case "fatal":
		logLevelValue = logrus.FatalLevel
	case "error":
		logLevelValue = logrus.ErrorLevel
	case "warn":
		logLevelValue = logrus.WarnLevel
	case "info":
		logLevelValue = logrus.InfoLevel
	case "debug":
		logLevelValue = logrus.DebugLevel
	case "trace":
		logLevelValue = logrus.TraceLevel
	default:
		return fmt.Errorf("invalid log level: %s", logLevel)
	}

	logrus.SetLevel(logLevelValue)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	for _, line := range strings.Split(sentinelAsciiArt, "\n") {
		logrus.Infoln(line)
	}

	return nil
}

func initELClient(ctx context.Context, elClientConfig *elclient.Configuration) (*ethclient.Client, error) {
	elClient, err := ethclient.DialContext(ctx, elClientConfig.RPCURL)

	if err != nil {
		return nil, err
	}

	return elClient, nil
}

func initBlocksListenerService(blocksListenerConfig *listener.BlocksListenerConfiguration, elClient *ethclient.Client) listener.BlocksListener {
	blocksListenerService := listener.NewBlocksListener(blocksListenerConfig, elClient)

	return blocksListenerService
}

func initABIs() (abi.ABI, abi.ABI, error) {
	metaMorphoFactoryABI, err := abi.JSON(strings.NewReader(bindings.MetaMorphoFactoryABI))
	if err != nil {
		return abi.ABI{}, abi.ABI{}, err
	}

	metaMorphoVaultABI, err := abi.JSON(strings.NewReader(bindings.MetaMorphoVaultABI))
	if err != nil {
		return abi.ABI{}, abi.ABI{}, err
	}

	return metaMorphoFactoryABI, metaMorphoVaultABI, nil
}

func initEventsListenerService(
	eventsListenerConfig *listener.EventsListenerConfiguration,
	elClient *ethclient.Client,
	metaMorphoFactoryABI abi.ABI,
	metaMorphoVaultABI abi.ABI,
) (listener.EventsListener, error) {
	eventsListenerService, err := listener.NewEventsListener(
		eventsListenerConfig,
		elClient,
		metaMorphoFactoryABI,
		metaMorphoVaultABI,
	)
	if err != nil {
		return nil, err
	}

	return eventsListenerService, nil
}

func initDecoderService(elClient *ethclient.Client, metaMorphoVaultABI abi.ABI) (decoder.Service, error) {
	decoderService, err := decoder.NewService(elClient, metaMorphoVaultABI)
	if err != nil {
		return nil, err
	}

	return decoderService, nil
}

func initIndexerService() indexer.Service {
	indexerService := indexer.NewService()

	return indexerService
}

func initNotifierService() notifier.Service {
	notificationConnectors := make([]notifier.NotificationConnector, 0)
	notificationConnectors = append(notificationConnectors, notifier.NewTerminalLogNotificationConnector())

	notifierService := notifier.NewService(notificationConnectors)

	return notifierService
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var configurationFilePath string

func init() {
	cobra.OnInitialize(initConfig)

	// Registering specific root flags.
	rootCmd.PersistentFlags().StringVarP(&configurationFilePath, "config", "c", "", "path to the configuration file")
}

func initConfig() {
	if configurationFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configurationFilePath)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".sentinel-config" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".sentinel-config")
	}

	viper.SetEnvPrefix("SENTINEL")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logrus.
			WithFields(logrus.Fields{
				"configuration-file": viper.ConfigFileUsed(),
			}).
			Debugln()
	}
}
