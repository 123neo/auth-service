package repository

import (
	"context"

	"go.uber.org/zap"
)

func (repo *Repo) GetUserByEmail(ctx context.Context, email string) (int64, error) {
	sqlStatement := `
	SELECT count(email) FROM users WHERE email = $1`

	repo.log.Info("Get Email:", zap.Any("get user by email", email))

	rows, err := repo.Db.QueryContext(ctx, sqlStatement, email)

	repo.log.Info("Rows:", zap.Any("rows:", rows))

	if err != nil {
		return 0, err
	}

	var count int64

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}

	repo.log.Info("count of email query:", zap.Any("count:", count))

	return count, nil

}
