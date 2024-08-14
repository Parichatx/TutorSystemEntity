package entity

import ("gorm.io/gorm"
		"time")

type Enrollment struct {
	gorm.Model
	EnrollmentDate time.Time

	// UserId ทำหน้าที่เป็น FK
	UserId *uint
	User   User  `gorm:"foreignKey:UserId"`

	// UserId ทำหน้าที่เป็น FK
	CourseID *uint
	Course   Course  `gorm:"foreignKey:CourseID"`
}