package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"questions/pkg/models"
)

var (
	DB *gorm.DB
	err error
)

func init()  {
	//database.DB, err = gorm.Open("mysql", "root:1234@/gindb?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open("sqlite3", "questions.db")
	if err != nil {
		fmt.Println("Status: ", err)
	}
	//defer DB.Close()
	DB.Debug()
	DB.LogMode(true)
	DB.AutoMigrate(&models.User{}, &models.Question{}, &models.Answer{}, &models.Tag{})

}
