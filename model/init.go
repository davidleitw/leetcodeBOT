package model

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// test database
func SetDB(db *gorm.DB) {
	DB = db
	migrate()
}

// 正式環境的資料庫
func ConnectionDatabase() {
	//key := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	key := fmt.Sprintf("port=%s host=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_PASSWORD"))

	//db, err := gorm.Open(mysql.Open(key), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(key), &gorm.Config{})
	if err != nil {
		log.Println("DB connect error = ", err.Error())
	}

	DB = db
	//CreateLeetCodeProblemsTable()
	migrate()
}

func migrate() {
	_ = DB.AutoMigrate(&Problem{})
	_ = DB.AutoMigrate(&Report{})
	_ = DB.AutoMigrate(&StudyGroup{})
}

func ClearTestDatabase() {
	_ = DB.Migrator().DropTable(&Report{})
	_ = DB.Migrator().DropTable(&StudyGroup{})
	migrate()
}
