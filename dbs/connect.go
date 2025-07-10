package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/citizen_issue")
	if err != nil {
		log.Fatal("Something went wrong:", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}
}
