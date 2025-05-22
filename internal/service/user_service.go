package service

import (
	"context"

	"github.com/Cladkoewka/grind-tracker/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByTelegramID(ctx context.Context, telegram_id int64) (*domain.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterOrGetUser(ctx context.Context, telegramID int64, username string) (*domain.User, error) {
	user, err := s.repo.GetByTelegramID(ctx, telegramID)
	if err == nil {
		return user, nil
	}

	newUser := &domain.User {TelegramID: telegramID, Username: username}
	err = s.repo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}