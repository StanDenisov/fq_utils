package profile

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID   uint
	SureName string
	LastName string
	Name     string
}
