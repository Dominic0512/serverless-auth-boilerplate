package config

import (
	"log"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"ENV"`
	DBDriver    string `mapstructure:"DB_DRIVER"`
	DBUsername  string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASS"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBName      string `mapstructure:"DB_NAME"`
}

func NewConfig() (*Config, error) {
	config := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
		return nil, err
	}

	return &config, nil
}

var ProviderSet = wire.NewSet(NewConfig)
