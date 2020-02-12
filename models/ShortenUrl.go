package models

import(
	"pudroid/database"
	"github.com/jinzhu/gorm"
	//"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ShortenUrl struct {
	gorm.Model
	Code string ` json:"code"`
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
func (u *ShortenUrl) Create() {
	db,err := database.DBConn()
	defer db.Close()
	if err != nil {
        panic(err.Error())
	}
	db.Create(&u)

}
func (u *ShortenUrl) Update() {
	db,err := database.DBConn()
	defer db.Close()
	if err != nil {
        panic(err.Error())
	}
	db.Save(&u)

}
func GetShortenUrlByCode(code string) (*ShortenUrl,error) {
	db,err := database.DBConn()
	defer db.Close()
	if err != nil {
        panic(err.Error())
	}
	sUrl := ShortenUrl{}

	shit := db.Where("code = ?", code).First(&sUrl).Error

	if shit!=nil {
		if(gorm.IsRecordNotFoundError(shit) ){
			return &sUrl,nil
		}
		return &sUrl,shit
	} else{
		return &sUrl,nil
	}
}
func GetAllShortenUrl() (*[]ShortenUrl,[]error) {
	var urls []ShortenUrl
	db,err := database.DBConn()
	defer db.Close()
	if err != nil {
        panic(err.Error())
	}
	errors := db.Find(&urls).GetErrors()
	if errors!=nil {
	return &urls,errors
	} else{
	return &urls,nil
	}
}


