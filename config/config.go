package config

import (
	"errors"
	"fmt"

	"github.com/CienciaArgentina/go-backend-commons/config"
)

// Configuration for the User Profile resource
type Configuration struct {
	Server struct {
		Port uint32
		Host string
	}
	DB      config.Database
	Verbose bool `yaml:"-"`
}

// Validate verifies that the configuration is correct
func (cfg Configuration) Validate() error {
	if cfg.Server.Port == 0 || cfg.Server.Port > uint32(1<<32-1) {
		return fmt.Errorf("%d: invalid port", cfg.Server.Port)
	}
	if cfg.Server.Host == "" {
		return errors.New("Invalid host")
	}
	return nil
}
