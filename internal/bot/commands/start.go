package commands

import (
	"context"
	"fmt"

	"github.com/Cladkoewka/grind-tracker/internal/service"
	"gopkg.in/telebot.v3"
)

type StartCommand struct {
	UserService *service.UserService
}

func (cmd *StartCommand) Handle(c telebot.Context) error {
	user, err := cmd.UserService.RegisterOrGetUser(context.Background(), c.Sender().ID, c.Sender().Username)
	if err != nil {
		return c.Send("Ошибка при регистрации пользователя.")
	}

	msg := fmt.Sprintf("Привет, %s! Я помогу тебе отслеживать твои навыки.", user.Username)
	return c.Send(msg)
}