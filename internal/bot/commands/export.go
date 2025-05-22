package commands

import (
	"context"
	"encoding/csv"
	"fmt"
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
	user, err := cmd.UserService.RegisterOrGetUser(context.Background(), c.Sender().ID, c.Sender().Username)
	if err != nil {
		return c.Send("Ошибка при получении пользователя.")
	}

	activities, err := cmd.ActivityService.GetUserActivities(context.Background(), user.ID)
	if err != nil {
		return c.Send("Ошибка при получении активностей.")
	}

	fileName := fmt.Sprintf("export_%d.csv", user.ID)
	file, err := os.CreateTemp("", fileName)
	if err != nil {
		return c.Send("Не удалось создать временный файл.")
	}
	defer os.Remove(file.Name())
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"Title", "Description", "Type", "XP", "CreatedAt"})
	for _, a := range activities {
		writer.Write([]string{a.Title, a.Description, a.Type, strconv.FormatInt(a.XP, 10), a.CreatedAt.Format("2006-01-02 15:04:05")})
	}
	writer.Flush()

	doc := &telebot.Document{File: telebot.FromDisk(file.Name()), FileName: "activities.csv"}
	return c.Send(doc)
}
