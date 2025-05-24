package commands

import (
	"context"
	"log/slog"
	"strconv"
	"strings"

	"github.com/Cladkoewka/grind-tracker/internal/domain"
	"github.com/Cladkoewka/grind-tracker/internal/service"
	"gopkg.in/telebot.v3"
)

type AddActivityCommand struct {
	UserService     *service.UserService
	ActivityService *service.ActivityService
}

func (cmd *AddActivityCommand) Handle(c telebot.Context) error {
	slog.Info("Обработка /add_activity", slog.String("payload", c.Message().Payload), slog.Int64("telegram_id", c.Sender().ID))

	args := strings.Split(c.Message().Payload, ";")
	if len(args) < 5 {
		slog.Warn("Неверное количество аргументов в /add_activity", slog.Int("arg_len", len(args)))
		return c.Send(`❌ Неверный формат команды.

<b>Правильный формат:</b>
<code>/add_activity skill_id; type; title; description; xp</code>

<b>Пояснение:</b>
• <code>skill_id</code> — номер навыка (см. /skills)  
• <code>type</code> — тип действия (например: видео, статья, пет-проект)  
• <code>title</code> — краткое название  
• <code>description</code> — что именно ты сделал  
• <code>xp</code> — количество полученного опыта

<b>Пример:</b>
<code>/add_activity 2; Видео на ютуб; Алгоритмы Сортировки; Посмотрел видео о быстрой сортировке; 10</code>

💡 Используй <b>/skills</b>, чтобы узнать ID навыков.`, telebot.ModeHTML)
	}

	skillID, err := strconv.ParseInt(strings.TrimSpace(args[0]), 10, 64)
	if err != nil {
		slog.Warn("Ошибка парсинга skill_id", slog.String("error", err.Error()))
		return c.Send("⚠️ Некорректный skill_id. Убедись, что ты вводишь число.")
	}

	xp, err := strconv.ParseInt(strings.TrimSpace(args[4]), 10, 64)
	if err != nil {
		slog.Warn("Ошибка парсинга XP", slog.String("error", err.Error()))
		return c.Send("⚠️ Некорректный XP. Убедись, что ты вводишь число.")
	}

	user, err := cmd.UserService.RegisterOrGetUser(context.Background(), c.Sender().ID, c.Sender().Username)
	if err != nil {
		slog.Error("Ошибка получения пользователя", slog.String("error", err.Error()))
		return c.Send("❌ Ошибка при получении пользователя.")
	}

	input := domain.AddActivityInput{
		UserID:      user.ID,
		SkillID:     skillID,
		Type:        strings.TrimSpace(args[1]),
		Title:       strings.TrimSpace(args[2]),
		Description: strings.TrimSpace(args[3]),
		XP:          xp,
	}

	slog.Info("Добавление активности", slog.Any("input", input))

	err = cmd.ActivityService.AddActivity(context.Background(), input)
	if err != nil {
		slog.Error("Ошибка при добавлении активности", slog.String("error", err.Error()))
		return c.Send("❌ Произошла ошибка при добавлении активности.")
	}

	slog.Info("Активность успешно добавлена", slog.Int64("user_id", user.ID))

	return c.Send("✅ Активность успешно добавлена! Продолжай в том же духе 💪")
}
