package models

import(
	"pudroid/database"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
)
type User struct{
	database.User
}

func (u *User) Create() {
	db,err := database.DBConn()
	defer db.Close()
	if err != nil {
        panic(err.Error())
	}
	db.Create(&u)

}
func (u *User) Update() {
	db,err := database.DBConn()
	defer db.Close()
	if err != nil {
        panic(err.Error())
	}
	db.Save(&u)

}
func GetUser(data map[string]interface{}) (*User,error) {
	db,err := database.DBConn()
	defer db.Close()
	if err != nil {
        panic(err.Error())
	}
	user := User{}

	shit := db.Where(data).First(&user).Error
	return &user,shit
}
