package database 

import (
	"github.com/jinzhu/gorm"
)

type ShortenUrl struct {
	gorm.Model
	Code string ` json:"code"`
	Url string `json:"url" binding:"required"`
}
type User struct{
	gorm.Model
	Username string `json:"username"`
	Passwords string `json:"passwords"`
}

func DBMigrate() (*gorm.DB){
	db,err := DBConn()
	if err != nil {
        panic(err.Error())
	}
	db = db.AutoMigrate(&ShortenUrl{},&User{})
	return db
}