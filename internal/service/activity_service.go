package service

import (
	"context"

	"github.com/Cladkoewka/grind-tracker/internal/domain"
)

type ActivityRepository interface {
	Create(ctx context.Context, activity *domain.Activity) error
	GetByUserID(ctx context.Context, userID int64) ([]domain.Activity, error)
	GetUserSkillXP(ctx context.Context, userID int64) (map[int64]int64, error)
}

type ActivityService struct {
	repo ActivityRepository
}

func NewActivityService(repo ActivityRepository) *ActivityService {
	return &ActivityService{repo: repo}
}

func (s *ActivityService) AddActivity(ctx context.Context, input domain.AddActivityInput) error {
	activity := &domain.Activity{
		UserID:      input.UserID,
		SkillID:     input.SkillID,
		Type:        input.Type,
		Title:       input.Title,
		Description: input.Description,
		XP:          input.XP,
	}
	return s.repo.Create(ctx, activity)
}

func (s *ActivityService) GetUserActivities(ctx context.Context, userID int64) ([]domain.Activity, error) {
	return s.repo.GetByUserID(ctx, userID)
}
