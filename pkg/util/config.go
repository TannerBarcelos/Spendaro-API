package util

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func ReadAppConfig(appEnv *string) {

	viper.SetConfigName(*appEnv)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msgf("Error reading config file, %s", err)
		panic(err)
	}

	log.Log().Msgf("Config file loaded successfully using config file: %s", viper.ConfigFileUsed())
}

func GetConfigString(key string) string {
	return viper.GetString(key)
}
