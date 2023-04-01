package config_test

import (
	"github.com/NurfitraPujo/finance-tracker/config"
	"github.com/spf13/viper"
	"os"
	"strings"
	"testing"
)

func TestGetProjectRootDir(t *testing.T) {
	t.Run("Should return valid project root dir", func(t *testing.T) {
		isCi := os.Getenv("IS_CI")
		cwd, _ := os.Getwd()
		want := strings.Replace(cwd, "/config", "", -1)
		if isCi == "true" {
			want = strings.Replace(cwd, "/finance-tracker/config", "", -1)
		}

		got := config.GetProjectRootDir()

		if got != want {
			t.Errorf("Wrong project root dir, want %v but got %v", want, got)
		}
	})

	t.Run("Should return valid project root dir in CI", func(t *testing.T) {
		viper.Set("IS_CI", "true")

		cwd, _ := os.Getwd()
		want := strings.Replace(cwd, "/config", "", -1)
		got := config.GetProjectRootDir()

		if got != want {
			t.Errorf("Wrong project root dir, want %v but got %v", want, got)
		}
	})
}

func TestLoadConfig(t *testing.T) {
	t.Run("ENV should be loaded properly", func(t *testing.T) {
		config.LoadConfig()
		if viper.GetString("database.host") == "" {
			t.Error("Database config is not loaded")
		}
		if viper.GetString("redis.host") == "" {
			t.Error("Redis config is not loaded")
		}
	})
}
