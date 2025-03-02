package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Db         `yaml:"db"`
}

type Db struct {
	AppPort      string `yaml:"APP_PORT"`
	PostgresDB   string `yaml:"POSTGRES_DB"`
	PostgresUser string `yaml:"POSTGRES_USER"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

var (
	once sync.Once
	cfg  *Config
)

func Load() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		configPath := os.Getenv("CONFIG_PATH")
		if configPath == "" {
			log.Fatal("CONFIG_PATH is not set")
		}

		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			log.Fatalf("config file does not exist: %s", configPath)
		}

		var localCfg Config

		if err := cleanenv.ReadConfig(configPath, &localCfg); err != nil {
			log.Fatalf("cannot read config: %s", err)
		}
		cfg = &localCfg
	})
	return cfg
}
