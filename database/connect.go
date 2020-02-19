package database 

import (
	"github.com/jinzhu/gorm"
    "fmt"
    "pudroid/config"
)


func DBConn() (*gorm.DB,error) {
  
   dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=True",
		config.Config.DBUser,
		config.Config.DBPass,
		config.Config.DBDriver,
		config.Config.DBName,
		"utf8")

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
        panic(err.Error())
	}
	return db, err
}
