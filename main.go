<<<<<<< HEAD
package main

import (
	"github.com/SA-Tutorsystem/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// เปิดการเชื่อมต่อกับฐานข้อมูล
	db, err := gorm.Open(sqlite.Open("sa-tutorsystem.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// ทำการ AutoMigrate เพื่ออัปเดต schema
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Course{},
		&entity.TutorProfile{}, 
		&entity.LoginHistory{},
		&entity.CourseCategory{},
		&entity.Payment{},
		&entity.Review{},
		&entity.Enrollment{},
		&entity.UserRole{},
	)
	if err != nil {
		panic("failed to migrate database schema: " + err.Error())
	}

	// การดำเนินการอื่นๆ
}
=======
package main

import (
	"github.com/SA-Tutorsystem/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// เปิดการเชื่อมต่อกับฐานข้อมูล
	db, err := gorm.Open(sqlite.Open("sa-tutorsystem.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// ทำการ AutoMigrate เพื่ออัปเดต schema
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Course{},
		&entity.TutorProfile{}, 
		&entity.LoginHistory{},
		&entity.CourseCategory{},
		&entity.Payment{},
		&entity.Review{},
		&entity.Enrollment{},
		&entity.UserRole{},
	)
	if err != nil {
		panic("failed to migrate database schema: " + err.Error())
	}

	// การดำเนินการอื่นๆ
}
>>>>>>> c7cb4fa00cf71ed56c0c00f17444f88d87ba31da
