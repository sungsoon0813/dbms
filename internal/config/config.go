package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/dlintw/goconf"
)

var (
	configFile = flag.String("config", "/usr/local/etc/dbms.conf", "Absolute path of the configuration file")
)

type Config struct {
	BaseDir         string
}

func New() *Config {
	return &Config{}
}

func (c *Config) Read() error {
	readConfigData, err := goconf.ReadConfigFile(*configFile)
	if err != nil {
		return fmt.Errorf("failed to read the config file: %v", err)
	}

	if err := c.readDefaultConfig(readConfigData); err != nil {
		return err
	}

	return nil
}

func (c *Config) readDefaultConfig(readConfigData *goconf.ConfigFile) error {
	var err error

	c.BaseDir, err = readConfigData.GetString("default", "base_dir")
	if err != nil || len(c.BaseDir) == 0 {
		return errors.New("invalid BaseDir directory")
	}
	if c.BaseDir[0] != '/' {
		return errors.New("BaseDir directory should be specified as an absolute path")
	}

	return nil
}
