package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Cladkoewka/grind-tracker/internal/config"
	"github.com/Cladkoewka/grind-tracker/internal/logger"
	"github.com/Cladkoewka/grind-tracker/internal/repository"
	"github.com/Cladkoewka/grind-tracker/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	DB              *pgxpool.Pool
	UserService     *service.UserService
	SkillService    *service.SkillService
	ActivityService *service.ActivityService
}

func NewContainer() (*Container, error) {
	cfg := config.Load()

	slog.SetDefault(logger.New(cfg.Log.LOG_LEVEL))

	db, err := pgxpool.New(context.Background(), cfg.DB.DATABASE_URL)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	// Проверка подключения
	if err := db.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}

	userRepo := repository.NewUserRepository(db)
	skillRepo := repository.NewSkillRepository(db)
	activityRepo := repository.NewActivityRepository(db)

	userService := service.NewUserService(userRepo)
	skillService := service.NewSkillService(skillRepo, activityRepo)
	activityService := service.NewActivityService(activityRepo)

	return &Container{
		DB:              db,
		UserService:     userService,
		SkillService:    skillService,
		ActivityService: activityService,
	}, nil
}
