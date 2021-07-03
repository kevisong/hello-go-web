package server

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Config defeines server config
type Config struct {
	Port int `yaml:"port"`
}

var defaultPort = 80

// Check checks config
func (c *Config) Check() error {
	if c.Port < 0 {
		errMsg := fmt.Sprintf("server.port %d invalid", c.Port)
		logrus.Error(errMsg)
		return errors.New(errMsg)
	} else if c.Port == 0 {
		warnMsg := fmt.Sprintf("server.port %d not specified, using default port %d", c.Port, defaultPort)
		logrus.Warn(warnMsg)
	}
	return nil
}
