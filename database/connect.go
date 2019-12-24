package database 

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    "fmt"
)

func DBConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "MiPu_Itus"
    dbName := "pudroid"
    fmt.Printf("Using the story in %s.\n", dbUser+":"+dbPass+"@tcp(104.155.224.180)/"+dbName)
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(104.155.224.180)/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}