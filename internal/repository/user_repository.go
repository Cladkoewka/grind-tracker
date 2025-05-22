package repository

import (
	"context"

	"github.com/Cladkoewka/grind-tracker/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (telegram_id, username) VALUES ($1, $2) RETURNING id, created_at`
	return r.db.QueryRow(ctx, query, user.TelegramID, user.Username).
		Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) GetByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error) {
	query := `SELECT id, telegram_id, username, created_at FROM users WHERE telegram_id = $1`
	row := r.db.QueryRow(ctx, query, telegramID)

	var user domain.User
	err := row.Scan(&user.ID, &user.TelegramID, &user.Username, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}