package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	Birthday  time.Time

	// UserRoleID ทำหน้าที่เป็น FK
	UserRoleID *uint
	UserRole   UserRole `gorm:"foreignKey:UserRoleID"`

	// 1 user สร้างได้หลาย course
	Courses []Course `gorm:"foreignKey:UserId"`

	// 1 user สามารถมีหลาย login history
	LoginHistorys []LoginHistory `gorm:"foreignKey:UserId"`

	// 1 user สามารถมีหลาย enrollment
	Enrollments []Enrollment `gorm:"foreignKey:UserId"`

	// 1 user สามารถมีหลาย payment
	Payments []Payment `gorm:"foreignKey:UserId"`

	// 1 user สามารถมีหลาย review
	Reviews []Review `gorm:"foreignKey:UserId"`
}