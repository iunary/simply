package config

import (
	"log"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

const (
	ENVPREFIX = "SIMPLY"
)

func New(path string, logger *log.Logger) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigFile(path)
	v.SetEnvPrefix(ENVPREFIX)

	if err := v.ReadInConfig(); err != nil {
		log.Default().Println(err.Error())
		return nil, err
	}
	v.AutomaticEnv()
	logger.Printf("using %s config", path)
	return v, nil
}

var ProviderSet = wire.NewSet(New)
