package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~กรุณากรอกชื่อ"`
	Email      string
	CustomerID string
}
