package service

import (
	"context"
	"math"
	"sort"

	"github.com/Cladkoewka/grind-tracker/internal/domain"
)

type SkillRepository interface {
	GetAll(ctx context.Context) ([]domain.Skill, error)
	GetByID(ctx context.Context, id int64) (*domain.Skill, error)
}
type SkillService struct {
	repo SkillRepository
	activityRepo ActivityRepository
}

func NewSkillService(repo SkillRepository, activityRepo ActivityRepository) *SkillService {
	return &SkillService{repo: repo, activityRepo: activityRepo}
}

func (s *SkillService) ListSkills(ctx context.Context) ([]domain.Skill, error) {
	return s.repo.GetAll(ctx)
}

func (s *SkillService) GetUserSkillProgress(ctx context.Context, userID int64) ([]domain.SkillProgress, error) {
	xpBySkill, err := s.activityRepo.GetUserSkillXP(ctx, userID)
	if err != nil {
		return nil, err
	}

	skills, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []domain.SkillProgress
	for _, skill := range skills {
		xp := xpBySkill[skill.ID]
		if xp == 0 {
			continue
		}

		progress := domain.SkillProgress{
			SkillID:   skill.ID,
			SkillName: skill.Name,
			TotalXP:   xp,
			Level:     calculateLevel(xp),
		}
		result = append(result, progress)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].TotalXP > result[j].TotalXP
	})

	return result, nil
}

func calculateLevel(xp int64) int64 {
	var (
		baseXP   float64 = 100.0 
		exponent float64 = 1.5   
		level    int64   = 1
	)

	var totalXP float64 = 0

	for {
		requiredXP := baseXP * math.Pow(float64(level), exponent)
		if float64(xp) < totalXP+requiredXP {
			break
		}
		totalXP += requiredXP
		level++
	}

	return level
}
