package entity

import("gorm.io/gorm"
		"time") 

type LoginHistory struct {
	gorm.Model
	LoginTimestamp time.Time

	// UserId ทำหน้าที่เป็น FK
	UserId *uint
	User   User  `gorm:"foreignKey:UserId"`

}