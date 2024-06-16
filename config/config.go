package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
	GPTAPIKey     string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	viper.AutomaticEnv() // Read environment variables

	config := &Config{
		ServerAddress: viper.GetString("SERVER_ADDRESS"),
		DatabaseURL:   viper.GetString("DATABASE_URL"),
		GPTAPIKey:     viper.GetString("GPT_API_KEY"),
	}

	// Fallback to default values if not set
	if config.ServerAddress == "" {
		config.ServerAddress = "localhost:8080" // Default value
	}
	if config.DatabaseURL == "" {
		config.DatabaseURL = "root:password@tcp(127.0.0.1:3306)/dbname" // Default value
	}
	if config.GPTAPIKey == "" {
		config.GPTAPIKey = "default_gpt_api_key" // Default value
	}

	return config
}
