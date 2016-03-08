package kaonashi

import (
	"log"

	"github.com/BurntSushi/toml"
)

// AppConfig struct
type AppConfig struct {
	Debug        bool   `toml:"debug"`
	Testing      bool   `toml:"testing"`
	DatabasePath string `toml:"database_path"`
	ServerPort   string `toml:"server_port"`
}

// NewAppConfig creates new config
func NewAppConfig(configFilePath string) (*AppConfig, error) {
	var config AppConfig
	if _, err := toml.DecodeFile(configFilePath, &config); err != nil {
		log.Fatalf("failed to create AppConfig from file: %s", err)
		return nil, err
	}
	return &config, nil
}

// NewAppDefaultConfig create new default config
func NewAppDefaultConfig() *AppConfig {
	var config *AppConfig
	config = &AppConfig{
		Debug:        true,
		Testing:      true,
		ServerPort:   "8588",
		DatabasePath: "kaonashi.db",
	}
	return config
}
