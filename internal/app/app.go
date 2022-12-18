package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/wire"
	"github.com/iunary/simply/internal/transports/http"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Application struct {
	name       string
	logger     *log.Logger
	httpServer *http.Server
}

type Option func(app *Application) error

func HTTPServerOption(svr *http.Server) Option {
	return func(app *Application) error {
		app.logger.Println("application name is", app.name)
		app.httpServer = svr
		return nil
	}
}

func New(v *viper.Viper, logger *log.Logger, options ...Option) (*Application, error) {
	app := &Application{
		logger: logger,
		name:   v.GetString("app.name"),
	}

	for _, option := range options {
		if err := option(app); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (a *Application) Start() error {
	if a.httpServer != nil {
		if err := a.httpServer.Start(); err != nil {
			return errors.Wrap(err, "http server start error")
		}
	}

	return nil
}

func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	s := <-c
	a.logger.Printf("receive a signal %s", s.String())
	if a.httpServer != nil {
		if err := a.httpServer.Stop(); err != nil {
			a.logger.Printf("stop http server error %s", err.Error())
		}
	}

	os.Exit(0)
}

var ProviderSet = wire.NewSet(New)
