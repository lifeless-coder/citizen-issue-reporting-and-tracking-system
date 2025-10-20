package main

import (
	"citizen_issue/dbs"
	"citizen_issue/internal/user"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	dbs.ConnectDB()
	user.UserMain()
	defer dbs.CloseDB()
	r := mux.NewRouter()
	lg := &user.Loginreq{}
	r.HandleFunc("user/login", lg.Login).Methods("GET")

}
