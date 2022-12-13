package repository

import (
	"auth-service/models"
	"context"
	"database/sql"

	"go.uber.org/zap"
)

type Repo struct {
	Db  *sql.DB
	log *zap.Logger
}

func NewRepo(Db *sql.DB, log *zap.Logger) Repository {
	return &Repo{
		Db:  Db,
		log: log,
	}
}

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (int64, error)
}

func (repo *Repo) CreateUser(ctx context.Context, user *models.User) error {
	sqlStatement := `
	INSERT INTO users (user_id, first_name, last_name, email, contact, password)
	VALUES ($1, $2, $3, $4, $5, $6)`

	repo.log.Info("User:", zap.Any("user struct", user))

	_, err := repo.Db.ExecContext(ctx, sqlStatement, user.ID, user.FirstName, user.LastName, user.Email, user.Contact, user.Password)
	if err != nil {
		return err
	}

	return nil
}
