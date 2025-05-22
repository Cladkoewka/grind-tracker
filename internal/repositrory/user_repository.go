package repositrory

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

