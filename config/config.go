package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var Cfg *Config

type Config struct {
	Logger   *LoggerConfig   `yaml:"logger"`
	Server   *ServerConfig   `yaml:"server"`
	Services *ServicesConfig `yaml:"services"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

type ServerConfig struct {
	Address string `yaml:"address"`
}

func Init(p string) error {
	barr, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}

	Cfg = &Config{}
	err = yaml.Unmarshal(barr, Cfg)
	if err != nil {
		return err
	}
	return nil
}
