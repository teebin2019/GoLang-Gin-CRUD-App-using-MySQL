package main

import (
	"GoCRUDApplicationMySQL/app/models"
	"GoCRUDApplicationMySQL/app/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:1234@tcp(localhost:3306)/vegetable_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// ตรวจสอบว่า column มีอยู่จริงก่อนลบ
	if db.Migrator().HasColumn(&models.User{}, "password") {
		db.Migrator().DropColumn(&models.User{}, "password")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Vegetable{})

	// Get the underlying *sql.DB from GORM
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRouter(sqlDB)
	r.Run(":3000")
}
