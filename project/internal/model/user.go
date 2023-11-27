package model

import (
	//"project/internal/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string `json:"name"  gorm:"unique"`
	Email        string `json:"email"  gorm:"unique"`
	PasswordHash string `json:"-" validate:"required"`
	Dob          string `json:"dob" validate:"required" `
}

type UserSignup struct {
	UserName string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Dob      string `json:"dob" validate:"required" `
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserForgetPassword struct {
	Email string `json:"email" validate:"required"`
	DOB   string `json:"dob" validate:"required"`
}

type UserUpdatePassword struct {
	Email          string `json:"email" validate:"required"`
	Password       string `json:"password" validate:"required"`
	RetypePassword string `json:"retype_password" validate:"required"`
	Otp            string `json:"otp" validate:"required"`
}
