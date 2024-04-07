package util

import (
	"testing"

	"github.com/spf13/viper"
)

func TestReadAppConfig(t *testing.T) {
	t.Run("Reads_Config", func(t *testing.T) {
		appEnv := "development"
		ReadAppConfig(&appEnv)

		if viper.ConfigFileUsed() == "" {
			t.Errorf("Failed to load config file")
		}

		if GetConfigString("env") != "development" {
			t.Errorf("Failed to get config value")
		}
	})

	t.Run("Fails_Reading_Config", func(t *testing.T) {
		appEnv := "nonexistentfile"

		// Test that the code panics when the config file is not found
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		ReadAppConfig(&appEnv)
	})
}
