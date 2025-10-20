package user

import (
	"citizen_issue/dbs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func reInputPass(hashPass string) {
	var input string
	fmt.Scanln(&input)
	if !dbs.ComparePass(hashPass, input) {
		log.Println("wrong password")
		reInputPass(hashPass)
	} else {
		fmt.Println("login success")
	}

}

type loginreq struct {
	mobile   string `json:"mobile"`
	Password string `json:"password"`
}

func (req *loginreq) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hashPass string
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	exstMob, userId := dbs.IsNumberExist(req.mobile)
	if exstMob {
		hashPass, err = dbs.GetPassword(userId)
	} else {
		log.Fatal("userId not exist", err)
	}
	if !dbs.ComparePass(hashPass, req.Password) {
		log.Println("wrong password")
		reInputPass(hashPass)

	} else {
		fmt.Println("login success")
		UserDashboard(w, r)
	}

}
