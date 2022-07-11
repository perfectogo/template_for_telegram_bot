package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App string

	Environment string // development, staging, production

	LogLevel string // debug, info, warn, error, dpanic, panic, fatal
	HTTPPort string
	BasePath string

	TelegramBotApiToken string
	ChatUsername        string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "telegram_bot"))

	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "staging"))

	config.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	config.BasePath = cast.ToString(getOrReturnDefaultValue("BASE_PATH", "/v1"))

	config.TelegramBotApiToken = cast.ToString(getOrReturnDefaultValue("TELEGRAM_BOT_API_TOKEN", "5409049240:AAGAI-a0rUH_Y5nobNXbJl8Ym4xCipDW-XE"))
	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "rabbit"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "jasurbek"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1001"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
