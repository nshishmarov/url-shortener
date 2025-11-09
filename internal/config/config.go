package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Mode string `yaml:"mode" env-default:"dev"`
	Server ServerConfig `yaml:"server"`
	DBConfig DBConfig `yaml:"db"`
}

type ServerConfig struct {
	Port int `yaml:"port" env-default:"8080"`
	Host string `yaml:"host" env-default:"localhost"`
	Timeout time.Duration `yaml:"timeout" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DBConfig struct {
	DSN string `yaml:"dsn"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg
}