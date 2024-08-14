package entity

import ("gorm.io/gorm"
		"time")

type Payment struct {
	gorm.Model
	Amount float32
	DateTime time.Time

	// UserId ทำหน้าที่เป็น FK
	UserId *uint
	User   User  `gorm:"foreignKey:UserId"`

	// UserId ทำหน้าที่เป็น FK
	CourseID *uint
	Course   Course  `gorm:"foreignKey:CourseID"`
}