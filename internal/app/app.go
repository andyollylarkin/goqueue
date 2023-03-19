package app

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"goqueue/internal/configs"
	"goqueue/pkg/eventbus"
	"goqueue/pkg/utils"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	config configs.Config
	logger *logrus.Logger
}

func NewApplication(config configs.Config) Application {
	return Application{config: config, logger: utils.NewLogger()}
}

func (app *Application) Start() error {
	app.logger.Info("Application started")
	stopAppCh := make(chan os.Signal)
	signal.Notify(stopAppCh, unix.SIGKILL, syscall.SIGTERM)
	eventBus := eventbus.NewEventBus()
	<-stopAppCh
	eventShutdown := NewShutdownEvent()
	// emit new shutdown event and wait until all subscribers done work
	eventBus.EmitEvent(eventShutdown)
	eventShutdown.WaitUntilEventWillBeProcessed()
	return nil
}
