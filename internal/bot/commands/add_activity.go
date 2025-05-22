package commands

import (
	"context"
	"strconv"
	"strings"

	"github.com/Cladkoewka/grind-tracker/internal/service"
	"github.com/Cladkoewka/grind-tracker/internal/domain"
	"gopkg.in/telebot.v3"
)

type AddActivityCommand struct {
	UserService     *service.UserService
	ActivityService *service.ActivityService
}

func (cmd *AddActivityCommand) Handle(c telebot.Context) error {
	args := strings.Split(c.Message().Payload, ";")
	if len(args) < 5 {
		return c.Send("Формат: /add_activity <skill_id>; <type>; <title>; <description>; <xp>")
	}

	skillID, err := strconv.ParseInt(strings.TrimSpace(args[0]), 10, 64)
	if err != nil {
		return c.Send("Некорректный skill_id")
	}

	xp, err := strconv.ParseInt(strings.TrimSpace(args[4]), 10, 64)
	if err != nil {
		return c.Send("Некорректный XP")
	}

	user, err := cmd.UserService.RegisterOrGetUser(context.Background(), c.Sender().ID, c.Sender().Username)
	if err != nil {
		return c.Send("Ошибка при получении пользователя.")
	}

	input := domain.AddActivityInput{
		UserID:      user.ID,
		SkillID:     skillID,
		Type:        strings.TrimSpace(args[1]),
		Title:       strings.TrimSpace(args[2]),
		Description: strings.TrimSpace(args[3]),
		XP:          xp,
	}

	err = cmd.ActivityService.AddActivity(context.Background(), input)
	if err != nil {
		return c.Send("Ошибка при добавлении активности.")
	}

	return c.Send("Активность успешно добавлена.")
}
