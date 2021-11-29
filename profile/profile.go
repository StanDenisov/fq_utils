package profile

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	SureName string
	LastName string
	Name     string
}
