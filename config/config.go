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

type DB struct {
	Host     string `env-required:"true" yaml:"host"   env:"DB_HOST"`
	Port     string `env-required:"true" yaml:"port"   env:"DB_PORT"`
	Username string `env-required:"true" yaml:"username"   env:"DB_USERNAME"`
	DBName   string `env-required:"true" yaml:"dbname" env:"DB_NAME"`
	SSLMode  string `env-required:"true" yaml:"sslmode" env:"DB_SSLMODE"`
}

type Config struct {
	App  `yaml:"app"`
	HTTP `yaml:"http"`
	DB   `yaml:"db"`
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
