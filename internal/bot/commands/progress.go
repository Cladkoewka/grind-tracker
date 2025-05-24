package commands

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/Cladkoewka/grind-tracker/internal/service"
	"gopkg.in/telebot.v3"
)

type ProgressCommand struct {
	UserService  *service.UserService
	SkillService *service.SkillService
}

func (cmd *ProgressCommand) Handle(c telebot.Context) error {
	ctx := context.Background()
	telegramID := c.Sender().ID
	username := c.Sender().Username

	slog.Info("Обработка команды /progress", 
		slog.Int64("telegram_id", telegramID),
		slog.String("username", username),
	)

	user, err := cmd.UserService.RegisterOrGetUser(ctx, telegramID, username)
	if err != nil {
		slog.Error("Ошибка при получении пользователя", slog.String("error", err.Error()))
		return c.Send("Ошибка при получении пользователя.")
	}

	progressList, err := cmd.SkillService.GetUserSkillProgress(ctx, user.ID)
	if err != nil {
		slog.Error("Ошибка при получении прогресса пользователя", 
			slog.Int64("user_id", user.ID),
			slog.String("error", err.Error()),
		)
		return c.Send("Ошибка при получении прогресса.")
	}

	if len(progressList) == 0 {
		slog.Info("У пользователя нет активности", slog.Int64("user_id", user.ID))
		return c.Send("Вы ещё не добавили ни одной активности.")
	}

	var sb strings.Builder
	sb.WriteString("Ваш прогресс по навыкам:\n\n")
	for _, p := range progressList {
		sb.WriteString(fmt.Sprintf("🛠 %s — XP: %d, Уровень: %d\n", p.SkillName, p.TotalXP, p.Level))
	}

	slog.Info("Успешно получен прогресс пользователя", 
		slog.Int64("user_id", user.ID),
		slog.Int("skills_count", len(progressList)),
	)

	return c.Send(sb.String())
}
