package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func ConnectDB() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/citizen_issue")
	if err != nil {
		log.Fatal("Something went wrong:", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}

}
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
