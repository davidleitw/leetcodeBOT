package test

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	test_userID  string = "a001"
	test_guildID string = "guild001"
)

func getTestDB() *gorm.DB {
	key := fmt.Sprintf("port=%s host=%s user=%s dbname=%s sslmode=disable password=%s", "5432",
		"127.0.0.1", "postgres", "leetcodeDB", "postgres")
	db, _ := gorm.Open(postgres.Open(key), &gorm.Config{})
	return db
}
