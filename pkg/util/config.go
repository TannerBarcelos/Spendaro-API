package util

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func ReadAppConfig(appEnv *string) {

	viper.SetConfigName(*appEnv)        // name of config file (without extension)
	viper.AddConfigPath("config")       // path to look for the config file in
	viper.AddConfigPath("../../config") // path to look for the config file in (for tests)

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msgf("Error reading config file, %s", err)
		panic(err)
	}

	log.Log().Msgf("Config file loaded successfully. Using config file: %s", viper.ConfigFileUsed())
}

func Read(key string) string {
	return viper.GetString(key)
}
