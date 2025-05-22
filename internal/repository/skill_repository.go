package repository

import (
	"context"

	"github.com/Cladkoewka/grind-tracker/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SkillRepository struct {
	db *pgxpool.Pool
}

func NewSkillRepository(db *pgxpool.Pool) *SkillRepository {
	return &SkillRepository{db: db}
}

func (r *SkillRepository) GetAll(ctx context.Context) ([]domain.Skill, error) {
	query := `SELECT id, name, description FROM skills`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []domain.Skill
	for rows.Next() {
		var skill domain.Skill
		if err := rows.Scan(&skill.ID, &skill.Name, &skill.Description); err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}
	return skills, nil
}

func (r *SkillRepository) GetByID(ctx context.Context, id int64) (*domain.Skill, error) {
	query := `SELECT id, name, description FROM skills WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var skill domain.Skill
	if err := row.Scan(&skill.ID, &skill.Name, &skill.Description); err != nil {
		return nil, err
	}
	return &skill, nil
}