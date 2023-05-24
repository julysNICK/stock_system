package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DB_URL 					 string        `mapstructure:"DB_URL"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	HTTP_SERVER_ADDRESS         string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TOKEN_SYMMETRIC_KEY    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	ACCESS_TOKEN_DURATION    time.Duration        `mapstructure:"ACCESS_TOKEN_DURATION"` 
	REFRESH_TOKEN_DURATION    time.Duration        `mapstructure:"REFRESH_TOKEN_DURATION"` 
}


func LoadConfig(path string) (config Config, err error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return 
	}

	err = viper.Unmarshal(&config)

	return
}