package repository

import (
	"gorm.io/gorm"
)

// DatabaseRepo handles interactions with the database
type DatabaseRepo struct {
	Db *gorm.DB
}

// NewDatabaseRepo creates a new instance of DatabaseRepo
func NewDatabaseRepo(db *gorm.DB) *DatabaseRepo {
	return &DatabaseRepo{
		Db: db,
	}
}
