package service

import (
	"context"

	"github.com/Cladkoewka/grind-tracker/internal/domain"
)

type SkillRepository interface {
	GetAll(ctx context.Context) ([]domain.Skill, error)
	GetByID(ctx context.Context, id int64) (*domain.Skill, error)
}

type SkillService struct {
	repo SkillRepository
}

func NewSkillService(repo SkillRepository) *SkillService {
	return &SkillService{repo: repo}
}

func (s *SkillService) ListSkills(ctx context.Context) ([]domain.Skill, error) {
	return s.repo.GetAll(ctx)
}
