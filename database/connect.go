package database 

import (
	"github.com/jinzhu/gorm"
    "fmt"
)


func DBConn() (*gorm.DB,error) {
    dbDriver := "104.155.224.180"
    dbUser := "pudroid"
    dbPass := "MiPu_Itus"
    dbName := "pudroid"
   dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=True",
		dbUser,
		dbPass,
		dbDriver,
		dbName,
		"utf8")

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
        panic(err.Error())
	}
	return db, err
}
