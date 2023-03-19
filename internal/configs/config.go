package configs

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"net"
	"os"
)

type Config struct {
	Version string `yaml:"version"`
	Storage struct {
		Type string `yaml:"type"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	SecureConnection struct {
		UseTls         bool   `yaml:"use_tls"`
		CACertFilePath string `yaml:"ca_cert_file_path"`
		CertFilePath   string `yaml:"cert_file_path"`
		KeyFilePath    string `yaml:"key_file_path"`
	}
	CompressionType string `yaml:"compression_type"`
	ListenAddr      string `yaml:"listen"`
	ListenPort      uint16 `yaml:"port"`
}

func MustNewConfig(debugMode bool) Config {
	return mustPrepareConfig(debugMode)
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
	err = validateTlsConfig(config)
	if err != nil {
		panic(err)
	}
	err = validateNetworkAddr(&config)
	if err != nil {
		panic(err)
	}
	return config
}

func validateNetworkAddr(cfg *Config) error {
	listenToAddr(cfg)
	if ip := net.ParseIP(cfg.ListenAddr); ip == nil {
		return fmt.Errorf("IP address: %s invalid format", cfg.ListenAddr)
	}
	return nil
}

func listenToAddr(cfg *Config) {
	const AllAddr = "*"
	const LocalHost = "127.0.0.1"

	switch cfg.ListenAddr {
	case "":
		cfg.ListenAddr = LocalHost
	case AllAddr:
		cfg.ListenAddr = "0.0.0.0"
	default:
		return
	}
}

func validateTlsConfig(cfg Config) error {
	if cfg.SecureConnection.UseTls && cfg.SecureConnection.CertFilePath == "" {
		return errors.New("TLS mode enabled, but path to cert is empty")
	}
	if cfg.SecureConnection.UseTls && cfg.SecureConnection.KeyFilePath == "" {
		return errors.New("TLS mode enabled, but path to key is empty")
	}
	if cfg.SecureConnection.UseTls && cfg.SecureConnection.CACertFilePath == "" {
		return errors.New("TLS mode enabled, but path to CA cert is empty")
	}
	return nil
}
