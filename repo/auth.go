package repo

import (
	"avenger/entity"
)

func (dbh DBHandler) Register(user entity.User) (entity.User, error) {
	result := dbh.DB.Select("username", "password").Create(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	result = dbh.DB.Where("username = ?", user.Username).First(&user)
	return user, nil
}
