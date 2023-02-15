package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~กรุณากรอกชื่อ"`
	Email      string
	CustomerID string `valid:"required~รหัสลูกค้าไม่ถูกต้อง ,matches(^([L|M|H])([0-9]{7}$))~รูปแบบเบอร์โทรผู้รับไม่ถูกต้อง"`
}
