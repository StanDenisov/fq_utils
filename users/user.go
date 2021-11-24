package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string
	Phone     string
	Password  string
	ProfileId uint
}
