package infrastructure

import (
	"context"
	"github.com/bordunosp/ddd/example/app/user/domain"
	"github.com/google/uuid"
	"log"
)

func NewUserService(log *log.Logger) (domain.IUserService, error) {
	return &userService{
		logger: log,
	}, nil
}

type userService struct {
	logger *log.Logger
}

func (u *userService) SendCreatedEmail(ctx context.Context, userId uuid.UUID) error {
	//TODO implement me
	u.logger.Println("SendCreatedEmail called with userId: " + userId.String())
	return nil
}
