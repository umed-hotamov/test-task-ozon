package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
}

type PostgresConfig struct {
	Database string
	User     string
	Password string
}

type ServerConfig struct {
	port string
}

func GetConfig() *Config {
	v, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config %v", err)
	}

	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("Error parsing config %v", err)
	}

	return cfg
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigType("yml")
	v.SetConfigName("config")
	v.AddConfigPath("./config/")

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found")
		}
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &cfg, nil
}
