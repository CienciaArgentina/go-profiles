package config

import (
	"errors"
	"fmt"
)

// TODO: Remove me
// Configuration for the User Profile resource
type Configuration struct {
	Server struct {
		Port uint32
		Host string
	}
	DB      Database
	Verbose bool `yaml:"-"`
}

// Database is the DB configuration struct extending config.DataDatabase
type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Protocol string `yaml:"protocol"`
	Port     int    `yaml:"port"`
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
