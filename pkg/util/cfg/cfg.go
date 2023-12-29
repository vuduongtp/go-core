package cfgutil

import (
	"github.com/caarlos0/env/v5"
	"github.com/joho/godotenv"
)

// Load loads configuration from local .env file
func Load(out interface{}, stage string) error {
	if err := LoadLocalENV(stage); err != nil {
		return err
	}

	if err := env.Parse(out); err != nil {
		return err
	}

	return nil
}

// LoadConfig loads configuration from .env file
func LoadConfig(out interface{}, appName, stage string) error {

	return Load(out, stage)
}

// LoadLocalENV reads .env* files and sets the values to os ENV
func LoadLocalENV(stage string) error {
	basePath := ""
	if stage == "test" {
		basePath = "testdata/"
	}
	// // local config per stage
	// if stage != "" {
	// 	godotenv.Load(basePath + ".env." + stage + ".local")
	// }

	// local config
	godotenv.Load(basePath + ".env.local")

	// // per stage config
	// if stage != "" {
	// 	godotenv.Load(basePath + ".env." + stage)
	// }

	// default config
	return godotenv.Load(basePath + ".env")
}
