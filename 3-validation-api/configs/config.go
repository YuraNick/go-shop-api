package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Mail MailConfig
}

type MailConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, default config use")
	}
	return &Config{
		Mail: MailConfig{
			Email:    os.Getenv("Email"),
			Password: os.Getenv("Password"),
			Address:  os.Getenv("Address"),
		},
	}
}
