package users

import (
	"github.com/StanDenisov/fq_utils/profile"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Phone    string
	Password string
	Profile  profile.Profile
}
