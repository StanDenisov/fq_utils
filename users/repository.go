package users

import (
	"errors"

	"github.com/StanDenisov/fq_utils/profile"
	"gorm.io/gorm"
)

func GetUserIDByLoginAndPassword(db gorm.DB, u User) (uint, error) {
	user := User{}
	err := db.Where("phone = ? AND password = ?", u.Phone, u.Password).Find(&user)
	if err.Error != nil {
		return 0, errors.New("Error is %s" + err.Error.Error())
	}
	return user.ID, nil
}

func CreateUserAndProfile(db gorm.DB, u User) {
	err := db.Create(&User{Phone: u.Phone, Password: u.Password,
		Profile: profile.Profile{
			Model:    gorm.Model{},
			SureName: u.Profile.SureName,
			LastName: u.Profile.LastName,
			Name:     u.Profile.Name,
		}})
	if err.Error != nil {
		panic(err.Error.Error())
	}
}
