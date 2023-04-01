package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

func LoadConfig() {
	rootDir := GetProjectRootDir()

	env := os.Getenv("APP_ENV")
	if env != "" {
		viper.Set("APP_ENV", env)
	} else {
		viper.SetDefault("APP_ENV", "development")
	}

	ci := os.Getenv("ON_CI")
	if ci != "" {
		viper.Set("ON_CI", "true")
	} else {
		viper.SetDefault("ON_CI", "false")
	}

	envName := "env-" + viper.GetString("APP_ENV")

	viper.SetConfigName(envName)
	viper.SetConfigType("json")
	viper.AddConfigPath(rootDir + "/config")
	_ = viper.ReadInConfig()

	if !IsProduction() {
		viper.Debug()
	}
}

func GetProjectRootDir() string {
	workDir, _ := os.Getwd()
	dirArr := strings.Split(workDir, "/")

	rootDirs := make([]string, len(dirArr))
	// Handle duplicate folder name in ci eg: Github action checkout
	ciDuplicate := false
	for i, dir := range dirArr {
		if IsCI() {
			if dir == "finance-tracker" {
				if !ciDuplicate {
					ciDuplicate = true
					rootDirs[i] = dir
				} else {
					rootDirs[i] = dir
					break
				}
			} else {
				if ciDuplicate {
					break
				} else {
					rootDirs[i] = dir
					continue
				}
			}
		} else {
			if dir == "finance-tracker" {
				if !ciDuplicate {
					ciDuplicate = true
					rootDirs[i] = dir
				} else {
					break
				}
			} else {
				if ciDuplicate {
					break
				} else {
					rootDirs[i] = dir
					continue
				}
			}
		}
	}

	var finalPath = ""
	for _, path := range rootDirs {
		if path == "" {
			continue
		}
		finalPath = finalPath + "/" + path
	}
	return finalPath
}

func IsDevelopment() bool {
	return viper.GetString("APP_ENV") == "development"
}

func IsProduction() bool {
	return viper.GetString("APP_ENV") == "production"
}

func IsTest() bool {
	return viper.GetString("APP_ENV") == "test"
}

func IsCI() bool {
	return viper.GetString("IS_CI") == "true"
}
