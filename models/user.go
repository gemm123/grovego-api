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
	Routes    []Route `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	ImagePath string `json:"imagePath"`
}
