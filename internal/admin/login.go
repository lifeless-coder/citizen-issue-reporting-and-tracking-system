package admin

import (
	"citizen_issue/dbs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginReq struct {
	email string `json:"email"`
	pass  string `json:"password"`
}

func (a *LoginReq) Login(w http.ResponseWriter, r *http.Request) int {
	w.Header().Set("Content-Type", "application/json")
	var err error
	var hashPass string
	err = json.NewDecoder(r.Body).Decode(&a)
	exstMob, userId := dbs.IsAdminMailExist(a.email)
	if exstMob {
		hashPass, err = dbs.GetAdminPassword(userId)
	} else {
		log.Fatal("userId not exist", err)
	}
	if !dbs.ComparePass(hashPass, a.pass) {
		log.Println("wrong password, try again")
		a.Login(w, r)
		fmt.Println("try again later")
		return 0

	} else {
		fmt.Println("login success")
		return userId
	}
	return userId
}
