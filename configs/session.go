package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type SessiConfig struct {
	SecretKey string
}

func SessionConfig() SessiConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	config := SessiConfig{
		SecretKey: viper.GetString("session.secretKey"),
	}

	return config
}

func GetSessionSecretKey() string {
	return SessionConfig().SecretKey
}
