package database

import (
	"github.com/KEVISONG/hello-go-web/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Store defiens db storage
type Store struct {
	db *gorm.DB
}

// NewStore factory
func NewStore(dsn string) (dbStore *Store, err error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return &Store{db: db}, nil
}
