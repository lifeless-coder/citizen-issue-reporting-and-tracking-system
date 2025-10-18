package main

import (
	"citizen_issue/dbs"
	"citizen_issue/internal/user"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbs.ConnectDB()
	user.UserMain()
	defer dbs.CloseDB()
}
