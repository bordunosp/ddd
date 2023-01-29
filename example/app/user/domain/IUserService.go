package domain

import (
	"context"
	"github.com/google/uuid"
)

type IUserService interface {
	SendCreatedEmail(cnt context.Context, userId uuid.UUID) error
}
