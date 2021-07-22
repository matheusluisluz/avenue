package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ServerConfiguration struct {
	Port string
}

type Configuration struct {
	Server ServerConfiguration
	Upload string
}

func Execute() Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration Configuration
	fmt.Println("configuration.Server.Port", configuration.Server.Port)
	fmt.Println("configuration.Upload", configuration.Upload)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return configuration
}
