package database

import (
	"github.com/KEVISONG/hello-go-web/pkg/models"
)

// GetUserByID GetUserByID
func (s *Store) GetUserByID(id uint) (models.User, error) {
	user := models.User{}
	db := s.db.Where("id = ?", id).Find(&user)
	if db.Error != nil {
		return user, db.Error
	}
	return user, nil
}

// PostUser PostUser
func (s *Store) PostUser(u models.User) (id uint, err error) {
	db := s.db.Create(&u)
	if db.Error != nil {
		return 0, db.Error
	}
	return u.ID, nil
}
