package dao

import "gorm.io/gorm"

// Account represents an account dao
type Account struct {
	gorm.Model
	Balance float64
}
