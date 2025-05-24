package commands

import (
	"context"
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/Cladkoewka/grind-tracker/internal/service"
	"gopkg.in/telebot.v3"
)

type ExportCommand struct {
	UserService     *service.UserService
	ActivityService *service.ActivityService
}

func (cmd *ExportCommand) Handle(c telebot.Context) error {
	slog.Info("Обработка /export", slog.Int64("telegram_id", c.Sender().ID))

	user, err := cmd.UserService.RegisterOrGetUser(context.Background(), c.Sender().ID, c.Sender().Username)
	if err != nil {
		slog.Error("Ошибка при получении пользователя", slog.String("error", err.Error()))
		return c.Send("Ошибка при получении пользователя.")
	}

	activities, err := cmd.ActivityService.GetUserActivities(context.Background(), user.ID)
	if err != nil {
		slog.Error("Ошибка при получении активностей", slog.String("error", err.Error()))
		return c.Send("Ошибка при получении активностей.")
	}

	fileName := fmt.Sprintf("export_%d.csv", user.ID)
	file, err := os.CreateTemp("", fileName)
	if err != nil {
		slog.Error("Не удалось создать временный файл", slog.String("error", err.Error()))
		return c.Send("Не удалось создать временный файл.")
	}
	defer os.Remove(file.Name())
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"Title", "Description", "Type", "XP", "CreatedAt"})
	for _, a := range activities {
		writer.Write([]string{
			a.Title,
			a.Description,
			a.Type,
			strconv.FormatInt(a.XP, 10),
			a.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	writer.Flush()

	slog.Info("Файл экспортирован", slog.String("path", file.Name()), slog.Int("activity_count", len(activities)))

	doc := &telebot.Document{File: telebot.FromDisk(file.Name()), FileName: "activities.csv"}
	return c.Send(doc)
}
