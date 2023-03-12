package configs

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Version string `yaml:"version"`
	Storage struct {
		Type string `yaml:"type"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
}

func New(debugMode bool) Config {
	mustPrepareConfig(debugMode)
	return Config{}
}

func mustPrepareConfig(debugMode bool) Config {
	var namePart string = ""
	if debugMode {
		namePart = ".example"
	}
	configName := fmt.Sprintf("../../config%s.yaml", namePart)
	cfgFile, err := os.ReadFile(configName)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		panic(err)
	}
	config := Config{}
	err = yaml.Unmarshal(cfgFile, &config)
	if err != nil {
		panic(err)
	}
	return config
}
