package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Account is our model, which corresponds to the "accounts" table
type Account struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CommonID     uuid.UUID `gorm:"type:uuid;index:idx_common_id;not null;"`
	ClerkUserID  string    `gorm:"unique;index:idx_clerk_user_id;not null;"`
	FirstName    string    `gorm:"not null;"`
	LastName     string    `gorm:"not null;"`
	UserName     string    `gorm:"unique;index:idx_user_name;index:idx_user_name_email;where:user_name IS NOT NUll;"`
	EmailAddress string    `gorm:"unique;index:idx_email;index:idx_user_name_email;not null;"`
}
