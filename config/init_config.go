package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var config *viper.Viper = viper.New()

func InitConfig() *viper.Viper {
	_, filename, _, _ := runtime.Caller(0)
	// Move to the root directory by going up one or more levels from the current file
	rootPath := filepath.Join(filepath.Dir(filename), "..")

	// Set the configuration file explicitly to the root directory
	config.SetConfigFile(filepath.Join(rootPath, ".env"))

	// Alternatively, you can add the root path to search paths
	config.AddConfigPath(rootPath)

	// Read the config file
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return config
}
