package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServer     `yaml:"http_server"`
	DatabaseConfig `yaml:"database_config"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env-default:"localhost"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"potgres"`
	DBName   string `yaml:"dbname" env-default:"lexoread"`
	DBPort   string `yaml:"port" env-default:"5432"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

type HTTPServer struct {
	Port        string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	configPath := "config/config.yaml"
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	return &cfg
}

func (c *Config) GetStorageDSN() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host, c.User, c.Password, c.DBName, c.DBPort, c.SSLMode)
	return dsn
}
