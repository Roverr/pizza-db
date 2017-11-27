package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const appName = "PIZZA"

// DatabaseConfig contains database related configurations.
type DatabaseConfig struct {
	URL                    string        `envconfig:"DB_URL" required:"true" desc:"Database connection URL."`
	ConnectionTimeout      time.Duration `envconfig:"DB_CONN_TIMEOUT_DUR" default:"60s" desc:"Timeout after this duration if DB is not available."`
	ReconnectSleepInterval time.Duration `envconfig:"DB_RECONN_SLEEP_DUR" default:"1s" desc:"Sleep N duration before retrying to connect."`
}

// Config contains the configuration of this application.
type Config struct {
	ListenAddress string `envconfig:"LISTEN_ADDRESS" default:"127.0.0.1:8080" desc:"Listen on this address for incoming connections"`
	DataClean     bool   `envconfig:"DATA_CLEAN" default:"false" desc:"Indicates if the application is running in data clean mode"`
	DatabaseConfig
}

// Parse parses the configuration from the environment/command line flags
func Parse() (*Config, error) {
	cfg := Config{}
	if err := envconfig.Process(appName, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
