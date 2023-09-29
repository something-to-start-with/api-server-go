package main

import (
	"log"

	"github.com/something-to-start-with/api-server-go/internal/app/apiserver"
	"github.com/spf13/viper"
)

func main() {
	cfg := loadConfig()
	apiserver.Run(cfg)
}

func loadConfig() *apiserver.Config {
	viper.SetConfigName("api-server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	cfg := new(apiserver.Config)

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("Error unmarshal config file: %s", err)
	}

	return cfg
}
