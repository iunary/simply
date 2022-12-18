package main

import (
	"flag"
	"log"

	"github.com/iunary/simply/internal/app"
	"github.com/iunary/simply/internal/transports/http"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func newApp(v *viper.Viper, logger *log.Logger, hs *http.Server) (*app.Application, error) {
	a, err := app.New(v, logger, app.HTTPServerOption(hs))
	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	return a, nil
}

func main() {
	logger := log.Default()
	configFile := flag.String("config", "config/config.yml", "config file path")
	flag.Parse()
	app, err := CreateApp(*configFile, logger)
	if err != nil {
		panic(err)
	}
	logger.Println("starting application")
	if err := app.Start(); err != nil {
		panic(err)
	}
	app.AwaitSignal()
	logger.Println("closing application")
}
