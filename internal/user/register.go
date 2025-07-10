package user

import (
	"citizen_issue/dbs"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/mail"
	"regexp"
	"strings"
)

type User struct {
	usrname    string
	nid        string
	mobileNo   string
	email      string
	gender     string
	profession string
	birthday   string
	picture    string
	password   string
}

func checkNum(mb string) string { //return valid number
	var mobile string
	mb = strings.TrimSpace(mb)
	regex := "^01[3-9][0-9]{8}$"
	re := regexp.MustCompile(regex)
	valid := re.MatchString(mb)
	if !valid {
		fmt.Println("Invalid mobile number")
		fmt.Scanln(&mobile)
		return checkNum(mobile)
	}
	return mb

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

func checkEmail(ml string) string {
	var email string
	ml = strings.TrimSpace(ml)
	addr, err := mail.ParseAddress(ml)
	if err != nil {
		fmt.Println("invalid email", err)
		fmt.Scanln(&email)
		return checkEmail(email)
	}
	return addr.Address
}

func Register() {
	var u User
	fmt.Scanf("%s", &u.usrname)
	fmt.Scanln(&u.nid)
	fmt.Scanln(&u.mobileNo)
	u.mobileNo = checkNum(u.mobileNo)
	fmt.Scanln(&u.email)
	u.email = checkEmail(u.email)
	fmt.Scanln(&u.gender)
	fmt.Scanln(&u.profession)
	u.profession = strings.TrimSpace(u.profession)
	fmt.Scanln(&u.birthday)
	fmt.Scanln(&u.picture)

	dbs.InsertIntoUser(u.usrname, u.nid, u.mobileNo, u.password, u.gender, u.profession, u.birthday, u.picture, u.email)
}
