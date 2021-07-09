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
