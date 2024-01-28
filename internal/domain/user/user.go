package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `db:"id"`
	DisplayName string    `db:"display_name"`
	Email       string    `db:"email"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Version     int       `db:"version"`
}

type CreateUserRequest struct {
	DisplayName, Email string
}

type UpdateUserRequest struct {
	ID                 uuid.UUID
	DisplayName, Email string
}
