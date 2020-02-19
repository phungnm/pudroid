package models

import(
	"pudroid/database"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type ShortenUrl struct {
database.ShortenUrl
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
func GetShortenUrl(data map[string]interface{}) (*ShortenUrl,error) {
	db,err := database.DBConn()
	defer db.Close()
	if err != nil {
        panic(err.Error())
	}
	sUrl := ShortenUrl{}

	shit := db.Where(data).First(&sUrl).Error

	if shit!=nil {
		
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


