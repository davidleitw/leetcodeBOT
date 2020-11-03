package model

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	key := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(key), &gorm.Config{})
	if err != nil {
		log.Println("DB connect error = ", err.Error())
	}
	DB = db
	migrate()
}

func migrate() {
	_ = DB.AutoMigrate(&Problem{})
}
