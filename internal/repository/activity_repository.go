package repository

import (
	"context"

	"github.com/Cladkoewka/grind-tracker/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ActivityRepository struct {
	db *pgxpool.Pool
}

func NewActivityRepository(db *pgxpool.Pool) *ActivityRepository {
	return &ActivityRepository{db: db}
}

func (r *ActivityRepository) Create(ctx context.Context, activity *domain.Activity) error {
	query := `
		INSERT INTO activities (user_id, skill_id, type, title, description, xp)
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, created_at
	`
	return r.db.QueryRow(ctx, query,
		activity.UserID,
		activity.SkillID,
		activity.Type,
		activity.Title,
		activity.Description,
		activity.XP,
	).Scan(&activity.ID, &activity.CreatedAt)
}

func (r *ActivityRepository) GetByUserID(ctx context.Context, userID int64) ([]domain.Activity, error) {
	query := `
		SELECT id, user_id, skill_id, type, title, description, xp, created_at
		FROM activities
		WHERE user_id = $1
		ORDER BY created_by DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []domain.Activity
	for rows.Next() {
		var a domain.Activity
		if err := rows.Scan(
			&a.ID, &a.UserID, &a.SkillID, &a.Type, &a.Title, &a.Description, &a.XP, &a.CreatedAt,
		); err != nil {
			return nil, err
		}
		activities = append(activities, a)
	}
	return activities, nil
}
