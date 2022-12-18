//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"github.com/iunary/simply/internal/app"
	"github.com/iunary/simply/internal/config"
	"github.com/iunary/simply/internal/controllers"
	"github.com/iunary/simply/internal/database"
	"github.com/iunary/simply/internal/repositories"
	"github.com/iunary/simply/internal/services"
	"github.com/iunary/simply/internal/transports/http"
	"log"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,
	controllers.ProviderSet,
	http.ProviderSet,
)

func CreateApp(cf string, logger *log.Logger) (*app.Application, error) {
	wire.Build(providerSet, newApp)
	return &app.Application{}, nil
}
