package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Host string
	Port string
}

func getConfigs() Config { 
	return Config{
		Port: ":" + os.Getenv("PORT"),		
	}
}

var DefaultConfig = getConfigs()
