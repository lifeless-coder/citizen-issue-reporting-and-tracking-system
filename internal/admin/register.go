package admin

import (
	"citizen_issue/dbs"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/mail"
	"strings"
)

type Admin struct {
	adminName string
	email     string
	password  string
	addedBy   int
}

func isValidEmail(ml string) string {
	var email string
	ml = strings.TrimSpace(ml)
	addr, err := mail.ParseAddress(ml)
	if err != nil {
		fmt.Println("invalid email", err)
		fmt.Scanln(&email)
		ml = isValidEmail(email)
	}
	return addr.Address
}
func hashPass(pass string) string {
	var newPass string
	val, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Println("something is wrong, re enter pass:", err)
		fmt.Scanln(&newPass)
		return hashPass(newPass)
	}
	return string(val)
}
func existEmail() string {
	var mail string
	fmt.Scanln(&mail)
	mail = isValidEmail(mail)
	dupliMail, _ := dbs.IsNidExist(mail)
	if dupliMail {
		log.Println("email exist enter again")
		return existEmail()
	}
	return mail
}

func (a *Admin) Register(ref int) {
	fmt.Scanln(&a.adminName)
	fmt.Scanln(&a.email)
	a.email = isValidEmail(a.email)
	exst, _ := dbs.IsAdminMailExist(a.email)
	if !exst {
		a.email = existEmail()
	}
	fmt.Scanf("%s", &a.password)
	a.password = hashPass(a.password)
	dbs.InsertIntoAdmin(a.adminName, a.email, a.password, ref)
}
