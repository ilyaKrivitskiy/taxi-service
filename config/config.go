package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
	Version string `env-required:"true" yaml:"version"    env:"APP_VERSION"`
}

type HTTP struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

type Log struct {
	Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
}

type DB struct {
	Username string `env-required:"true" yaml:"db_username"   env:"DB_USERNAME"`
	Host     string `env-required:"true" yaml:"db_host"   env:"DB_HOST"`
	Port     string `env-required:"true" yaml:"db_port"   env:"DB_PORT"`
	Name     string `env-required:"true" yaml:"db_name" env:"DB_NAME"`
	SslMode  string `env-required:"true" yaml:"db_sslmode" env:"DB_SSLMODE"`
}

type Config struct {
	App  `yaml:"app"`
	HTTP `yaml:"http"`
	Log  `yaml:"logger"`
	DB   `yaml:"databases"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
