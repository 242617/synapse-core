package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var Cfg *Config

type Config struct {
	Server *ServerConfig `yaml:"server"`
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
