package app

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	brokerModule "goqueue/internal/broker"
	"goqueue/internal/configs"
	"goqueue/pkg/eventbus"
	"goqueue/pkg/events"
	"goqueue/pkg/utils"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type Application struct {
	config configs.Config
	logger *logrus.Logger
}

func NewApplication(config configs.Config, isDebugMode bool) Application {
	return Application{config: config, logger: utils.NewLogger(isDebugMode)}
}

func (app *Application) Start() error {
	app.logger.Info("Application started")
	stopAppCh := make(chan os.Signal)
	signal.Notify(stopAppCh, unix.SIGKILL, syscall.SIGTERM)
	eventBus := eventbus.NewEventBus(app.logger)

	brokerConfig := brokerModule.BrokerCfg{ListenAddr: app.config.ListenAddr, ListenPort: strconv.Itoa(int(app.config.
		ListenPort))}

	broker := brokerModule.NewBroker(brokerConfig, app.logger)
	err := broker.Run()
	if err != nil {
		app.logger.Error(err)
	}

	<-stopAppCh
	eventShutdown := events.NewShutdownEvent()
	// emit new shutdown event and wait until all subscribers done work
	eventBus.EmitEvent(eventShutdown)
	app.logger.Info("Wait application shutdown")
	eventShutdown.WaitUntilEventWillBeProcessed()
	app.logger.Info("Application shutdown")
	return nil
}
