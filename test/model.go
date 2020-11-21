package test

import (
	"fmt"
	"log"

	"github.com/davidleitw/leetcodeBOT/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	key := fmt.Sprintf("root:@(localhost)/leetcodetest?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(key), &gorm.Config{})
	if err != nil {
		log.Println("DB connect error = ", err.Error())
	}
	DB = db
	model.CreateLeetCodeProblemsTable()
}
