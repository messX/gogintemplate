package config

import (
	"os"

	"github.com/messx/gogintemplate/infra/logger"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server ServerConfiguration
}

func SetupConfig(envFiles ...string) error {
	var configuration *Configuration
	// viper.SetConfigName("dev.env")
	// viper.AddConfigPath("./envs")
	if len(envFiles) != 0 {
		viper.SetConfigFile(envFiles[0])
	} else {
		logger.Debugf("COnf file: %s", os.Getenv("CONFIG_FILE_PATH"))
		// viper.SetConfigFile(os.Getenv("CONFIG_FILE_PATH"))
		viper.SetConfigFile("./envs/dev.env")
	}
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error to reading config file, %s", err)
		return err
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Errorf("error to decode, %v", err)
		return err
	}

	return nil
}
