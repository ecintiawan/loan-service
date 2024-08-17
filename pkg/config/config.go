package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ecintiawan/loan-service/pkg/env"
	"github.com/spf13/viper"
)

func NewConfig() *Config {
	var (
		config = &Config{}
	)

	readConfig(config)
	readSecrets(config)

	return config
}

func readConfig(config *Config) {
	var (
		name  = "loan-service"
		paths = []string{"/etc/loan-service/%s/", "./files/etc/loan-service/%s/"}
		env   = env.GetEnv()
	)

	viper.SetConfigType("json")
	viper.SetConfigName(name)
	for _, path := range paths {
		viper.AddConfigPath(fmt.Sprintf(path, env))
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	bConfig, err := json.Marshal(viper.AllSettings())
	if err != nil {
		log.Fatalf("error marshaling config: %v", err)
	}

	if err := json.Unmarshal(bConfig, config); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}
}

func readSecrets(secret *Config) {
	var (
		name  = "loan-service.secret"
		paths = []string{"/etc/credential/%s/", "./files/etc/credential/%s/"}
		env   = env.GetEnv()
	)

	viper.SetConfigType("json")
	viper.SetConfigName(name)
	for _, path := range paths {
		viper.AddConfigPath(fmt.Sprintf(path, env))
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading secret: %v", err)
	}

	bSecret, err := json.Marshal(viper.AllSettings())
	if err != nil {
		log.Fatalf("error marshaling config: %v", err)
	}

	if err := json.Unmarshal(bSecret, secret); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}
}
