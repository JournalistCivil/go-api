package mariadb

import (
	goapi "github.com/anon/go-api/user"
	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func (us *UserService) FindById(id string) (*goapi.User, error) {
	var u goapi.User
	row := us.Db.First(&u, id)
	if row.Error != nil {
		return nil, row.Error
	}

	return &u, nil
}
