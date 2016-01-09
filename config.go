package kaonashi

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/naoina/toml"
)

// AppConfig struct
type AppConfig struct {
	Debug        bool
	Testing      bool
	DatabasePath string
	ServerPort   string
}

// NewAppConfig creates new config
func NewAppConfig(configFilePath string) (*AppConfig, error) {
	f, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalf("failed to open config file: %s: %s", configFilePath, err)
		return nil, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("failed to read config file: %s", err)
		return nil, err
	}
	var config AppConfig
	if err := toml.Unmarshal(buf, &config); err != nil {
		log.Fatalf("failed to create AppConfig from file: %s", err)
		return nil, err
	}
	return &config, nil
}

// NewDefaultConfig create new default config
func NewDefaultConfig() *AppConfig {
	var config *AppConfig
	config = &AppConfig{
		Debug:        true,
		Testing:      true,
		ServerPort:   "8080",
		DatabasePath: "kaonashi.db",
	}
	return config
}
