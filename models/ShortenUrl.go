package models

import(
	"pudroid/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ShortenUrl struct {
	gorm.Model
	Code string `gorm:"unique" json:"code"`
	Url string `json:"url"`
}
func DBMigrate() (*gorm.DB){
	db,err := database.DBConn()
	if err != nil {
        panic(err.Error())
	}
	db = db.AutoMigrate(&ShortenUrl{})
	return db
}
func (u *ShortenUrl) Save() {
	db,err := database.DBConn()
	defer db.Close()
	if err != nil {
        panic(err.Error())
	}
	db.Save(&u)
}
func GetShortenUrlByCode(code string) (*ShortenUrl,[]error) {
	db,err := database.DBConn()
	defer db.Close()
	if err != nil {
        panic(err.Error())
	}
	sUrl := ShortenUrl{}
	errors := db.Where("code = ?", code).First(&sUrl).GetErrors()
	return &sUrl,errors
}
func GetAllShortenUrl() (*[]ShortenUrl,[]error) {
	var urls []ShortenUrl
	db,err := database.DBConn()
	defer db.Close()
	if err != nil {
        panic(err.Error())
	}
	errors := db.Find(&urls).GetErrors()
	return &urls,errors
	}


