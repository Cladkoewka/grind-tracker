package commands

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/Cladkoewka/grind-tracker/internal/service"
	"gopkg.in/telebot.v3"
)

type SkillsCommand struct {
	SkillService *service.SkillService
}

func (cmd *SkillsCommand) Handle(c telebot.Context) error {
	slog.Info("Обработка /skills", slog.Int64("telegram_id", c.Sender().ID))

	skills, err := cmd.SkillService.ListSkills(context.Background())
	if err != nil {
		slog.Error("Не удалось получить список навыков", slog.String("error", err.Error()))
		return c.Send("❌ Не удалось получить список навыков.")
	}

	if len(skills) == 0 {
		slog.Warn("Список навыков пуст")
		return c.Send("ℹ️ Список навыков пуст.")
	}

	var sb strings.Builder
	sb.WriteString("📚 <b>Доступные навыки:</b>\n\n")

	for _, skill := range skills {
		sb.WriteString(fmt.Sprintf("🛠 <b>%s</b>  (id: %d)\n", skill.Name, skill.ID))
		if strings.TrimSpace(skill.Description) != "" {
			sb.WriteString(fmt.Sprintf("  %s\n", skill.Description))
		}
		sb.WriteString("\n")
	}

	slog.Info("Навыки успешно отправлены", slog.Int("count", len(skills)))
	return c.Send(sb.String(), telebot.ModeHTML)
}
