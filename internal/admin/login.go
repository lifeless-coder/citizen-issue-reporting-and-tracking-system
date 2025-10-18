package admin

import (
	"citizen_issue/dbs"
	"fmt"
	"log"
)

func reInputPass(hashPass string, id int) int {

	return id
}
func (a *Admin) Login() int {
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
		log.Println("wrong password, try again")
		for i := 0; i < 9; i++ {
			var input string
			fmt.Scanln(&input)
			if !dbs.ComparePass(hashPass, input) {
				log.Println("wrong password, try again")
			} else {
				fmt.Println("login success")
				return userId
			}
		}
		fmt.Println("try again later")
		return 0

	} else {
		fmt.Println("login success")
		return userId
	}
	return userId
}
