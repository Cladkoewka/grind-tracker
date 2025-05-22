package bot

import (
	"github.com/Cladkoewka/grind-tracker/internal/bot/commands"
	"github.com/Cladkoewka/grind-tracker/internal/service"
	"gopkg.in/telebot.v3"
)

type CommandRouter struct {
	Bot             *telebot.Bot
	UserService     *service.UserService
	SkillService    *service.SkillService
	ActivityService *service.ActivityService
}

func NewRouter(bot *telebot.Bot, userService *service.UserService, skillService *service.SkillService, activityService *service.ActivityService) *CommandRouter {
	return &CommandRouter{
		Bot:             bot,
		UserService:     userService,
		SkillService:    skillService,
		ActivityService: activityService,
	}
}

func (r *CommandRouter) RegisterHandlers() {
	start := &commands.StartCommand{UserService: r.UserService}
	skills := &commands.SkillsCommand{SkillService: r.SkillService}
	addActivity := &commands.AddActivityCommand{
		UserService:     r.UserService,
		ActivityService: r.ActivityService,
	}
	export := &commands.ExportCommand{
		UserService:     r.UserService,
		ActivityService: r.ActivityService,
	}

	r.Bot.Handle("/start", start.Handle)
	r.Bot.Handle("/skills", skills.Handle)
	r.Bot.Handle("/add_activity", addActivity.Handle)
	r.Bot.Handle("/export", export.Handle)
}
