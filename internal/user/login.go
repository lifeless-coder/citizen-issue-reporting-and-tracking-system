package user

import (
	"citizen_issue/dbs"
	"encoding/json"
	"log"
	"net/http"
)

type Loginreq struct {
	mobile   string `json:"mobile"`
	Password string `json:"password"`
}

func (req *Loginreq) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hashPass string
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	exstMob, userId := dbs.IsNumberExist(req.mobile)
	if exstMob {
		hashPass, err = dbs.GetPassword(userId)
	} else {
		log.Fatal("userId not exist", err)
	}
	if !dbs.ComparePass(hashPass, req.Password) {
		log.Println("wrong password")
		req.Login(w, r)

	} else {
		log.Println("login success")
		UserDashboard(w, r)
	}

}
