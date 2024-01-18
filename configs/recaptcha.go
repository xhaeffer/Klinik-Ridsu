package configs

import (
	"fmt"

	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/spf13/viper"
)

type ChaptchaConfig struct {
    RecaptchaKey string
}

func RecaptchaConfig() ChaptchaConfig {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("fatal error config file: %s", err))
    }

    config := ChaptchaConfig{
        RecaptchaKey: viper.GetString("recaptcha.secretKey"),
    }

    recaptcha.Init(config.RecaptchaKey)

    return config
}