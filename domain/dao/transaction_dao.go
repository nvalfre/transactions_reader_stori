package dao

import (
	"gorm.io/gorm"
)

// Transaction represents a transaction_service dao
type Transaction struct {
	gorm.Model
	AccountID uint
	Date      string
	Amount    float64
	IsCredit  bool
}
