package user

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
	}

}
func Login() {
	var moblie string
	var pass string
	var hashPass string
	var err error
	fmt.Scanln(&moblie)
	fmt.Scanln(&pass)
	exstMob, userId := dbs.IsNumberExist(moblie)
	if exstMob {
		hashPass, err = dbs.GetPassword(userId)
	} else {
		log.Fatal("userId not exist", err)
	}
	if !dbs.ComparePass(hashPass, pass) {
		log.Println("wrong password")
		reInputPass(hashPass)

	} else {
		fmt.Println("login success")
	}

}
