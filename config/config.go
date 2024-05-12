package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppEnv string

const (
	appPrd   = "prd"
	appTest  = "test"
	appDev   = "dev"
	appLocal = "local"
)

type Config struct {
	Stage   AppEnv `env:STAGE, required`
	AppName string `env:APP_NAME`
}

func NewLoadConfig() (*Config, error) {
	stage := os.Getenv("STAGE")
	if stage == appLocal || stage == appTest {
		err := godotenv.Load("../local.env")
		if err != nil {
			// 失敗した場合はログを吐いてプログラムを終了
			log.Fatal("Error loading local.env file")
		}
	} else {
		err := godotenv.Load("../.env")
		if err != nil {
			// 失敗した場合はログを吐いてプログラムを終了
			log.Fatal("Error loading .env file")
		}
	}

	cfg := &Config{
		Stage:   "local",
		AppName: "atlas",
	}
	return cfg, nil
}

func (cfg *Config) IsLocal() bool {
	return cfg.Stage == "local" || cfg.Stage == "test"
}
