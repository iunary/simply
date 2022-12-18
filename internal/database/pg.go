package database

import (
	"context"
	"log"

	"github.com/google/wire"
	"github.com/iunary/simply/internal/entities/ent"
	"github.com/iunary/simply/internal/entities/ent/migrate"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Options struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

func NewOptions(v *viper.Viper, logger *log.Logger) (*Options, error) {
	options := new(Options)
	if err := v.UnmarshalKey("db", options); err != nil {
		return nil, errors.Wrap(err, "unmarshal db options error")
	}
	logger.Println("database options loaded successfully", options)
	return options, nil
}

func New(o *Options, logger *log.Logger) (*ent.Client, error) {
	logger.Println("connecting to database")

	opts := []ent.Option{ent.Log(logger.Println)}

	client, err := ent.Open(o.Driver, o.Source, opts...)
	if err != nil {
		logger.Fatalf("failed opening connection to db: %v", err)
	}
	if err := client.Schema.Create(context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		return nil, err
	}
	return client, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
