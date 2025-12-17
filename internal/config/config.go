package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Storage StorageConfig `mapstructure:"storage"`
	Persistence PersistenceConfig `mapstructure:"persistence"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}
type StorageConfig struct {}

type PersistenceConfig struct {
	Enabled bool `mapstructure:"enabled"`
	DataDir string `mapstructure:"data_dir"`
	Aof AofConfig `mapstructure:"aof"`
}

type AofConfig struct {
	Enabled bool `mapstructure:"enabled"`
	File string `mapstructure:"file"`
	SyncMode string `mapstructure:"sync_mode"`
	FlushInterval time.Duration `mapstructure:"flush_interval"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/echo-db")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("ECHO_DB")

	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", 6380)
	viper.SetDefault("persistence.enabled", false)
	viper.SetDefault("persistence.data_dir", "./data")
	viper.SetDefault("persistence.aof.enabled", false)
	viper.SetDefault("persistence.aof.file", "appendonly.aof")
	viper.SetDefault("persistence.aof.sync_mode", "everysec")
	viper.SetDefault("persistence.aof.flush_interval", "1s")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	if config.Persistence.Enabled && config.Persistence.DataDir != "" {
		if err := os.MkdirAll(config.Persistence.DataDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create data directory: %w", err)
		}
	}

	return &config, nil
}

