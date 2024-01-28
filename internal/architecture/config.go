package architecture

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port     int `yaml:"port" default:"8080"`
	Database struct {
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		Database     string `yaml:"database"`
		Schema       string `yaml:"schema"`
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		MigrationSrc string `yaml:"migration_src"`
	} `yaml:"db"`
}

func LoadConfig(path string) (Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
