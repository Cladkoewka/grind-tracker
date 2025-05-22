package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/Cladkoewka/grind-tracker/internal/service"
	"gopkg.in/telebot.v3"
)

type SkillsCommand struct {
	SkillService *service.SkillService
}

func (cmd *SkillsCommand) Handle(c telebot.Context) error {
	skills, err := cmd.SkillService.ListSkills(context.Background())
	if err != nil {
		return c.Send("Не удалось получить список навыков.")
	}

	if len(skills) == 0 {
		return c.Send("Список навыков пуст.")
	}

	var sb strings.Builder
	sb.WriteString("Доступные навыки:\n")
	for _, skill := range skills {
		sb.WriteString(fmt.Sprintf("- %s: %s\n", skill.Name, skill.Description))
	}

	return c.Send(sb.String())
}
