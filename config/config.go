package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Logger   LoggerConfig   `yaml:"logger"`
	DB       DBConfig       `yaml:"db"`
	Berth    BerthConfig    `yaml:"berth"`
	Server   ServerConfig   `yaml:"server"`
	Services ServicesConfig `yaml:"services"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

type DBConfig struct{}

type BerthConfig struct {
	Address string `yaml:"address"`
}

type ServerConfig struct {
	Address string `yaml:"address"`
}

func Init(p string) (*Config, error) {
	barr, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	if err = yaml.Unmarshal(barr, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
