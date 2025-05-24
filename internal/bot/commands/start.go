package commands

import (
	"context"
	"fmt"

	"github.com/Cladkoewka/grind-tracker/internal/service"
	"gopkg.in/telebot.v3"
)

var (
	Menu           = &telebot.ReplyMarkup{}
	BtnAbout       = Menu.Data("❓ О боте", "about_btn")
	BtnAddActivity = Menu.Data("➕ Добавить активность", "add_activity_btn")
	BtnSkills      = Menu.Data("📚 Список навыков", "skills_btn")
	BtnProgress    = Menu.Data("📈 Прогресс", "progress_btn")
	BtnExport      = Menu.Data("📤 Экспорт", "export_btn")
)

type StartCommand struct {
	UserService *service.UserService
}

func (cmd *StartCommand) Handle(c telebot.Context) error {
	user, err := cmd.UserService.RegisterOrGetUser(context.Background(), c.Sender().ID, c.Sender().Username)
	if err != nil {
		return c.Send("Ошибка при регистрации пользователя.")
	}

	msg := fmt.Sprintf("👋 Привет, %s! Я помогу тебе отслеживать твои навыки.\nВыбери действие:", user.Username)

	Menu.Inline(
		Menu.Row(BtnAbout),
		Menu.Row(BtnAddActivity),
		Menu.Row(BtnSkills),
		Menu.Row(BtnProgress),
		Menu.Row(BtnExport),
	)

	return c.Send(msg, Menu)
}
