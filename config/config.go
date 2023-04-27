package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type FileConfig struct {
	UploadPath string
	LogPath    string
	Env        string
}

type Config struct {
	DbConfig
	ApiConfig
	FileConfig
}

func (c *Config) ReadConfigFile() error {
	err := godotenv.Load(".env")
	if err != nil {
		return errors.New("failed to load .env file")
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	c.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"),
		ApiPort: os.Getenv("API_PORT"),
	}

	c.FileConfig = FileConfig{
		Env:        os.Getenv("ENV"),
		LogPath:    os.Getenv("LOG_PATH"),
		UploadPath: os.Getenv("UPLOAD_PATH"),
	}

	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" ||
		c.DbConfig.User == "" || c.DbConfig.Password == "" || c.ApiConfig.ApiHost == "" ||
		c.ApiConfig.ApiPort == "" || c.FileConfig.Env == "" || c.FileConfig.LogPath == "" ||
		c.FileConfig.UploadPath == "" {
		return errors.New("missing required environment variables")
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfigFile()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
