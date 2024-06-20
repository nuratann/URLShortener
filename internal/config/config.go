package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string     `yaml:"env" env-required:"true"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HttpServer  HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	confPath := os.Getenv("CONFIG_PATH")
	if confPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		log.Fatal("config file in CONFIG_PATH does not exist")
	}

	var conf Config

	if err := cleanenv.ReadConfig(confPath, &conf); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &conf
}
