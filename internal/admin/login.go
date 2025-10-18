package admin

import (
	"citizen_issue/dbs"
	"fmt"
	"log"
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
func (a *Admin) Login() {
	var err error
	var hashPass string
	fmt.Scanln(&a.email)
	fmt.Scanln(&a.password)
	exstMob, userId := dbs.IsAdminMailExist(a.email)
	if exstMob {
		hashPass, err = dbs.GetAdminPassword(userId)
	} else {
		log.Fatal("userId not exist", err)
	}
	if !dbs.ComparePass(hashPass, a.password) {
		log.Println("wrong password")
		reInputPass(hashPass)

	} else {
		fmt.Println("login success")
	}

}
