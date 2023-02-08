package main

import (
	"ginChat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:Jx131400@tcp(182.92.181.18:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.Message{}, &models.Contact{})
	// db.AutoMigrate(&models.UserBasic{})

	// Create
	// user := &models.UserBasic{}
	// user.Name = "test"
	// db.Create(user)
	// fmt.Println(user)

	// Uptate
	// db.Model(user).Update("PassWord", "1234")
}
