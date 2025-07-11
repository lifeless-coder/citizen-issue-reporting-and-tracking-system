package main

import (
	"citizen_issue/dbs"
	"citizen_issue/internal/user"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbs.ConnectDB()
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		user.Register()
	} else {
		user.Login()
	}
	defer dbs.CloseDB()
}
