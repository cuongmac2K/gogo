package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

// DBConfig config for DB
type DBConfig struct {
	Host string `yaml:"host"`
	Name string `yaml:"name"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Port string `yaml:"port"`
}

type Config struct {
	DBconf DBConfig `yaml:"db"`
}

func ReadConfig(configPath string) (*Config, error) {

	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
