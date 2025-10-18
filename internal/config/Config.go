package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Version  string   `yaml:"version"`
	Servers  Server   `yaml:"servers"`
	Redis    Redis    `yaml:"redis"`
	Database Database `yaml:"database"`
}
type Server struct {
	Port string `yaml:"port"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DataName string `yaml:"database"`
}

func Load(file string) (*Config, error) {
	readConfig, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	if err = yaml.Unmarshal(readConfig, config); err != nil {
		return nil, err
	}
	return config, nil
}
