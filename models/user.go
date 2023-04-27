package models

import (
	"time"

	"github.com/google/uuid"
)

type RegisterUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    int    `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string
	Username  string
	Email     string
	Phone     int
	Password  string
	ImagePath string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
