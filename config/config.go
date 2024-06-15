package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
	GPTAPIKey     string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return &Config{
		ServerAddress: viper.GetString("SERVER_ADDRESS"),
		DatabaseURL:   viper.GetString("DATABASE_URL"),
		GPTAPIKey:     viper.GetString("GPT_API_KEY"),
	}
}
