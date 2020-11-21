package test

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	test_userID  string = "a001"
	test_guildID string = "guild001"
)

func getTestDB() *gorm.DB {
	key := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, _ := gorm.Open(mysql.Open(key), &gorm.Config{})
	return db
}
