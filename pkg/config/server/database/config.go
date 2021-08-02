package database

import (
	"errors"

	"github.com/sirupsen/logrus"
)

// Config defines database config
type Config struct {
	DSN string `yaml:"dsn"`
}

func (c *Config) Check() error {

	if c.DSN == "" {
		errMsg := "database.dsn is empty"
		logrus.Error(errMsg)
		return errors.New(errMsg)
	}

	return nil

}
