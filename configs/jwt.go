package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type TokenConfig struct {
	SecretKey string
}

var conf TokenConfig

func JWTConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := viper.UnmarshalKey("jwt", &conf); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
}

func GetJWTSecretKey() []byte {
	return []byte(conf.SecretKey)
}
