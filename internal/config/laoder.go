package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig() *AppConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config: %v", err)
		panic(err)
	}

	log.Printf("Config file used: %s", viper.ConfigFileUsed())

	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	log.Printf("Loaded config: %+v", cfg)

	return &cfg
}
