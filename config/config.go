package config

import "github.com/spf13/viper"

func LoadConfig() {
	viper.SetDefault("APP_ENV", "development")

	currEnv := viper.GetString("APP_ENV")
	envName := "env." + currEnv
	viper.SetConfigName(envName)
	viper.SetConfigType("json")
}
