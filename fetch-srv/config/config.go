package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var (
	Cfg Config
)

type (
	Config struct {
		Env         string    `mapstructure:"env"`
		Port        string    `mapstructure:"port"`
		JWT         JWTConfig `mapstructure:"jwt"`
		FixerAPIKey string    `mapstructure:"fixerApiKey"`
	}

	JWTConfig struct {
		Secret         string        `mapstructure:"secret"`
		AccessTokenExp time.Duration `mapstructure:"accessTokenExp"`
	}
)

func Init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to initialize config, error: %v \n", err)
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("failed to unmarshal config, error: %v \n", err)
	}
}
