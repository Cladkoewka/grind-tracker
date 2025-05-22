package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB  DBConfig
	Bot BotConfig
	Log LogConfig
}

type DBConfig struct {
	DATABASE_URL string
}

type BotConfig struct {
	BOT_TOKEN string
}

type LogConfig struct {
	LOG_LEVEL string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		DB: DBConfig{
			DATABASE_URL: getEnv("DATABASE_URL", ""),
		},
		Bot: BotConfig{
			BOT_TOKEN: getEnv("BOT_TOKEN", ""),
		},
		Log: LogConfig{
			LOG_LEVEL: getEnv("LOG_LEVEL", "INFO"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}
