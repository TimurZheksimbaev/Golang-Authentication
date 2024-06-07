package config

import (
	"errors"
	"time"
	"github.com/spf13/viper"
)

type AppConfig struct {
	DatabaseHost string `mapstructure:"DB_HOST"`
	DatabaseUsername string `mapstructure:"DB_USERNAME"`
	DatabaseName string `mapstructure:"DB_NAME"`
	DatabasePort int `mapstructure:"DB_PORT"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort int `mapstructure:"SERVER_PORT"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRES_IN"`
	TokenMaxAge int `mapstructure:"TOKEN_MAX_AGE"`
	TokenSecret string `mapstructure:"TOKEN_SECRET"`
}

func LoadEnv() (*AppConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.New("Could not read config file")
	}
	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	return &appConfig, nil
}

