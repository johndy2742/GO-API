package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    string
	DBHost     string
	DBPort     string
}

func LoadConfig(configFilePath string) (Config, error) {
	var config Config

	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %v", err)
	}

	viper.AutomaticEnv() // Automatically read from environment variables

	config.DBUser = viper.GetString("PG_USER")
	config.DBPassword = viper.GetString("PG_PASSWORD")
	config.DBName = viper.GetString("PG_DBNAME")
	config.SSLMode = viper.GetString("PG_SSLMODE")
	config.DBHost = viper.GetString("PG_HOST")
	config.DBPort = viper.GetString("PG_PORT")

	return config, nil
}
