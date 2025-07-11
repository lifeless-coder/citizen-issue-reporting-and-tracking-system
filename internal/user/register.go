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

func isValidNum(mb string) string { //return valid number
	var mobile string
	mb = strings.TrimSpace(mb)
	regex := "^01[3-9][0-9]{8}$"
	re := regexp.MustCompile(regex)
	valid := re.MatchString(mb)
	if !valid {
		fmt.Println("Invalid mobile number")
		fmt.Scanln(&mobile)
		mb = isValidNum(mobile)
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
func existNum(mb string) string {
	fmt.Scanln(&mb)
	mb = isValidNum(mb)
	dupliNum, _ := dbs.IsNumberExist(mb)
	if dupliNum {
		return existNum(mb)
	}
	return mb
}
func existNid(nid string) string {
	fmt.Scanln(&nid)
	// nid = isValidNid(nid)
	dupliNid, _ := dbs.IsNidExist(nid)
	if dupliNid {
		log.Println("nid exist enter again")
		return existNid(nid)
	}
	return nid
}

//mobile nid already exist kina check korte hobe

func Register() {
	var u User
	fmt.Scanln(&u.usrname)
	fmt.Scanln(&u.nid)
	u.nid = strings.TrimSpace(u.nid)
	dupliNid, _ := dbs.IsNidExist(u.nid)
	if dupliNid {
		log.Println("nid exist enter again")
		u.nid = existNid(u.nid)
	}
	fmt.Scanln(&u.mobileNo)
	u.mobileNo = isValidNum(u.mobileNo)
	dupliNum, _ := dbs.IsNumberExist(u.mobileNo)
	if dupliNum {
		u.mobileNo = existNum(u.mobileNo)
	}
	fmt.Scanln(&u.email)
	u.email = isValidEmail(u.email)
	fmt.Scanln(&u.gender)
	fmt.Scanln(&u.profession)
	u.profession = strings.TrimSpace(u.profession)
	fmt.Scanln(&u.birthday)
	fmt.Scanln(&u.picture)
	fmt.Println("enter pass")
	fmt.Scanf("%s", &u.password)
	u.password = hashPass(u.password)

	dbs.InsertIntoUser(u.usrname, u.nid, u.mobileNo, u.password, u.gender, u.profession, u.birthday, u.picture, u.email)
}
