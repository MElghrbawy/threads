package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	ServerHost string
	ServerPort int `validate:"required,min=1,max=65535"`

	DBUser     string `validate:"required"`
	DBPassword string `validate:"required"`
	DBHost     string `validate:"required"`
	DBPort     int    `validate:"required,min=1,max=65535"`
	DBName     string `validate:"required"`
	DBSchema   string `validate:"required"`
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: no .env file found or unable to load")
	}

	config := &Config{
		ServerHost: getEnv("SERVER_HOST", ""),
		ServerPort: getEnvAsInt("SERVER_PORT", 8080),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     getEnvAsInt("DB_PORT", 5432),
		DBName:     getEnv("DB_NAME", ""),
		DBSchema:   getEnv("DB_SCHEMA", "public"),
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func ConstructDatabaseURL(c *Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?search_path=%s&sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSchema,
	)
}

func ConstructServerAddress(c *Config) string {
	return fmt.Sprintf("%s:%d", c.ServerHost, c.ServerPort)
}
