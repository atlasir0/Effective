package config

import (
	"flag"
	"os"
	"time"

	"log/slog"

	"github.com/spf13/viper"
)

type Config struct {
    Env      string     `mapstructure:"env"`
    Database Database   `mapstructure:"database"`
    HTTP     HTTPConfig `mapstructure:"http"`
}

type Database struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    Name     string `mapstructure:"name"`
}

type HTTPConfig struct {
    Port    int           `mapstructure:"port"`
    Timeout time.Duration `mapstructure:"timeout"`
}

func MustLoad() *Config {
    configPath := fetchConfigPath()
    if configPath == "" {
        panic("config path is empty")
    }

    return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        panic("config file does not exist: " + configPath)
    }

    viper.SetConfigFile(configPath)
    if err := viper.ReadInConfig(); err != nil {
        panic("cannot read config: " + err.Error())
    }

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        panic("cannot unmarshal config: " + err.Error())
    }

    slog.Debug("Loaded configuration", slog.Any("config", cfg))

    return &cfg
}

func fetchConfigPath() string {
    var res string

    flag.StringVar(&res, "config", "", "path to config file")
    flag.Parse()

    if res == "" {
        res = os.Getenv("CONFIG_PATH")
    }

    return res
}