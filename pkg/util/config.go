package util

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

func ReadAppConfig() {
	appEnv := flag.String("APP_ENV", "dev", "application runtime environment")
	flag.Parse()
	log.Printf("Application environment: %s", *appEnv)

	viper.SetConfigName(*appEnv)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s. Please provide a config file in the config directory", err)
		panic(err)
	}

	log.Printf("Using config file: %s", viper.ConfigFileUsed())
}

func GetConfigString(key string) string {
	return viper.GetString(key)
}
