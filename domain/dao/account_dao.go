package dao

import "gorm.io/gorm"

// Account represents an account dao
type Account struct {
	gorm.Model
	Name    string
	Email   string
	Balance float64
}
