package config

import (
	"fmt"
	"os"
	"resource-management/internal/lib/logger"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

const CONFIG_DIR string = "configs"

func init() {
	// config.ParseEnv: will parse env var in string value. eg: shell: ${SHELL}
	config.WithOptions(config.ParseEnv)

	// Add driver for support yaml content
	config.AddDriver(yaml.Driver)

	// Load default config
	err := config.LoadFiles(fmt.Sprintf("%s/default.yaml", CONFIG_DIR))
	if err != nil {
		logger.Fatal().Err(err).Msg("Error reading default config file")
	}

	// Get current environment (e.g., development, production)
	env := os.Getenv("APP_ENV")

	// Load env config
	if env != "" {
		err := config.LoadFiles(fmt.Sprintf("%s/%s.yaml", CONFIG_DIR, env))
		if err != nil {
			logger.Warn().Str("env", env).Err(err).Msg("Error reading config file")
		}
	}
}

func Get[K any](key string) (K, error) {
	var desiredConfig K
	err := config.BindStruct(key, &desiredConfig)

	return desiredConfig, err
}
