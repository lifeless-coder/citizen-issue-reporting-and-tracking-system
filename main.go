package main

import (
	"citizen_issue/dbs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbs.ConnectDB()
}
