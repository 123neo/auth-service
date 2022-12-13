package services

import (
	"context"

	"go.uber.org/zap"
)

func (s *service) GetUserByEmail(ctx context.Context, email string) (int64, error) {

	count, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return 0, err
	}

	s.log.Info("count of email query:", zap.Any("count:", count))

	return count, nil

}
