package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Address          string
	GracefulShutdown time.Duration
	DbUrl            string
}

// New creates a new config reference with default values.
func New() *Config {
	return &Config{
		Address:          "localhost:8080",
		GracefulShutdown: 10 * time.Second,
		DbUrl:            "file:certgen.db",
	}
}

// LoadFromEnv loads user defined configurations based on environment variables.
func (cfg *Config) LoadFromEnv() error {

	loadOptionalString("CERTGEN_ADDR", &cfg.Address)
	loadOptionalString("CERTGEN_DB_URL", &cfg.DbUrl)

	if err := loadOptionalDuration("CERTGEN_GRACEFUL_SHUTDOWN", &cfg.GracefulShutdown); err != nil {
		return err
	}

	return nil
}

func loadOptionalString(key string, targetValue *string) {
	if value, ok := os.LookupEnv(key); ok {
		*targetValue = value
	}
}

// func loadRequiredString(key string, targetValue *string) error {
// 	value, found := os.LookupEnv(key)
// 	if !found {
// 		return fmt.Errorf("required environment variable %s is not defined", key)
// 	}
// 	*targetValue = value
// 	return nil
// }

func loadOptionalDuration(key string, targetValue *time.Duration) error {
	value, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return fmt.Errorf("%s environment variable parsing as duration failed: %w", key, err)
	}

	*targetValue = duration
	return nil
}
