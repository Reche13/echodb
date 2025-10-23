package config

import (
	"fmt"
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

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}
