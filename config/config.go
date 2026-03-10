package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	BinanceApiKey     string
	BinanceApiSecret  string
	OKXApiKey         string
	OKXApiSecret      string
	OKXPassphrase     string
	BitfinexApiKey    string
	BitfinexApiSecret string
	PionexApiKey      string
	PionexApiSecret   string
}

func Load() (*Config, error) {
	// Load environment variables from all files, raise error if all of the files are not found
	envfiles := []string{".env", path.Join(os.Getenv("HOME"), ".ccyrc")}
	envfilesLoaded := 0
	for _, envfile := range envfiles {
		if _, err := os.Stat(envfile); os.IsNotExist(err) {
			continue
		}
		godotenv.Load(envfile)
		envfilesLoaded++
	}
	if envfilesLoaded == 0 {
		return nil, fmt.Errorf("no environment files found, please create one of the following files: %s", strings.Join(envfiles, ", "))
	}

	return &Config{
		BinanceApiKey:     os.Getenv("BINANCE_API_KEY"),
		BinanceApiSecret:  os.Getenv("BINANCE_API_SECRET"),
		OKXApiKey:         os.Getenv("OKX_API_KEY"),
		OKXApiSecret:      os.Getenv("OKX_API_SECRET"),
		OKXPassphrase:     os.Getenv("OKX_PASSPHRASE"),
		BitfinexApiKey:    os.Getenv("BFX_API_KEY"),
		BitfinexApiSecret: os.Getenv("BFX_API_SECRET"),
		PionexApiKey:      os.Getenv("PIONEX_API_KEY"),
		PionexApiSecret:   os.Getenv("PIONEX_API_SECRET"),
	}, nil
}
