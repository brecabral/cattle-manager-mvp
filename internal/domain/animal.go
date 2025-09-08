package domain

import (
	"time"

	"github.com/google/uuid"
)

type Animal struct {
	ID          uuid.UUID
	ExternalID  string
	Specie      string
	Race        string
	DateOfBirth time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
