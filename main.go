package main

import (
	"fmt"
	"viper-config-sample/configuration"
)

func main() {

	configPath := "resources/config.json"

	config, err := configuration.GetConfig(configPath)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Println("*** DATABASE INFO ***")
	fmt.Printf("- Server Port: %d\n", config.ServerPort)
	fmt.Printf("- Log Level: %s\n", config.LogLevel)
	fmt.Printf("- Database Host: %s\n", config.DBHost)
	fmt.Printf("- Database Port: %d\n", config.DBPort)
	fmt.Printf("- Database Username: %s\n", config.DBUsername)
	fmt.Printf("- Database Password: %s\n", config.DBPassword)

}
