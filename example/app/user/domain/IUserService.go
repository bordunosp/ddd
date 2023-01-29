package domain

import (
	"context"
	"github.com/google/uuid"
)

type IUserService interface {
	SendCreatedEmail(ctx context.Context, userId uuid.UUID) error
}
