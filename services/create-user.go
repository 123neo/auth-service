package services

import (
	"auth-service/models"
	"auth-service/repository"
	"context"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type service struct {
	repo repository.Repository
	log  *zap.Logger
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return &service{repo: repo, log: log}
}

type Service interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	GetUserByEmail(ctx context.Context, email string) (int64, error)
}

func (s *service) CreateUser(ctx context.Context, user *models.User) (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		s.log.Warn("Error in generating uuid..")
	}
	id := uuid.String()
	user.ID = id

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return "", err
	}
	return id, nil

}
