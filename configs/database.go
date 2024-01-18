package configs

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DbName   string `mapstructure:"dbname"`
}

func DbConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(os.Getenv("CONFIG_PATH"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
	}
}

func ReadDatabaseConfigs() []DatabaseConfig {
	var configs []DatabaseConfig
	err := viper.UnmarshalKey("databases", &configs)
	if err != nil {
		log.Fatalf("Failed to unmarshal database configurations: %v", err)
	}
	return configs
}

func FindDatabaseConfig(configs []DatabaseConfig, dbName string) (DatabaseConfig, error) {
	for _, config := range configs {
		if strings.EqualFold(config.Name, dbName) {
			return config, nil
		}
	}
	return DatabaseConfig{}, fmt.Errorf("database configuration for %s not found", dbName)
}
