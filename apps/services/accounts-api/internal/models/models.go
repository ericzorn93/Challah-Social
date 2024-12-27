package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Account is our model, which corresponds to the "accounts" table
type Account struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName string
	LastName  string
	NickName  string `gorm:"unique;index:idx_nick_name"`
	UserName  string `gorm:"unique;index:idx_user_name;index:idx_user_name_email"`
	Email     string `gorm:"unique;index:idx_email;index:idx_user_name_email"`
}
