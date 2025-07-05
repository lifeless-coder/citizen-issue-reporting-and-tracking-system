package main

import (
	"citizen_issue/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	database.ConnectDB()
}
