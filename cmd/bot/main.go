package main

import (
	"log/slog"
	"time"

	"github.com/Cladkoewka/grind-tracker/internal/app"
	"github.com/Cladkoewka/grind-tracker/internal/bot"
	"github.com/Cladkoewka/grind-tracker/internal/config"
	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/telebot.v3"
)

func main() {
	cfg := config.Load()

	container, err := app.NewContainer()
	if err != nil {
		slog.Error("ошибка инициализации контейнера", slog.String("error", err.Error()))
		return
	}

	botSettings := telebot.Settings{
		Token:  cfg.Bot.BOT_TOKEN,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	telegramBot, err := telebot.NewBot(botSettings)
	if err != nil {
		slog.Error("ошибка создания бота", slog.String("error", err.Error()))
		return
	}

	router := bot.NewRouter(
		telegramBot,
		container.UserService,
		container.SkillService,
		container.ActivityService,
	)

	router.RegisterHandlers()

	slog.Info("Бот запущен")
	telegramBot.Start()
}
