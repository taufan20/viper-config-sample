package configuration

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configuration struct {
	ServerPort int    `mapstructure:"server_port"`
	LogLevel   string `mapstructure:"log_level"`
	DBHost     string `mapstructure:"db_host"`
	DBPort     int    `mapstructure:"db_port"`
	DBUsername string `mapstructure:"db_username"`
	DBPassword string `mapstructure:"db_password"`
}

func GetConfig(configPath string) (Configuration, error) {
	config := Configuration{}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("failed to read config file: %v", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return config, nil

}
