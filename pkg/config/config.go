package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/KEVISONG/hello-go-web/pkg/config/debug"
	"github.com/KEVISONG/hello-go-web/pkg/config/server"
	"github.com/KEVISONG/hello-go-web/pkg/log"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Config defines configuration
type Config struct {
	Server server.Config `yaml:"server"`
	Debug  debug.Config  `yaml:"debug"`
	Log    log.Config    `yaml:"log"`
}

const defaultConfigFile = "./config.yml"

// C is default static config
var C Config

// Check checks if config is valid
func (c *Config) Check() error {
	return c.Log.Check()
}

// ToString converts server config to string
func (c *Config) ToString() (string, error) {
	configBytes, err := yaml.Marshal(c)
	if err != nil {
		errMsg := fmt.Sprintf("yaml.Marshal(%v) failed, error: %s", *c, err)
		logrus.Error(errMsg)
		return "", errors.New(errMsg)
	}
	return string(configBytes), nil
}

// LoadConfig loads configuration from file
func LoadConfig(configFile string, config *Config) error {
	configContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		errMsg := fmt.Sprintf("ioutil.ReadFile(%s) failed, error: %s", configFile, err)
		logrus.Error(errMsg)
		return errors.New(errMsg)
	}

	err = yaml.Unmarshal(configContent, &config)
	if err != nil {
		errMsg := fmt.Sprintf("yaml.Unmarshal %s failed, error: %s", configFile, err)
		logrus.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

// InitConfig inits configuration
func InitConfig(configFile string) error {

	if configFile == "" {
		configFile = defaultConfigFile
	}

	C = Config{}

	err := LoadConfig(configFile, &C)
	if err != nil {
		errMsg := fmt.Sprintf("LoadConfig(%s) failed, error: %s", configFile, err)
		logrus.Error(errMsg)
		return errors.New(errMsg)
	}

	configStr, _ := C.ToString()
	logrus.Info(string(configStr))

	return C.Check()

}
