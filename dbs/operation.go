package dbs

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func InsertIntoUser(uname, nid, mobile, pass, gender, prof, bday, pic, mail string) {
	query := "INSERT INTO user" +
		"(username, nid, mobile_num, user_password, gender, profession, birthdate, picture, email)" +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, uname, nid, mobile, pass, gender, prof, bday, pic, mail)
	if err != nil {
		log.Println("Error inserting user: ", err)
	}
}

func ComparePass(hashpass, givenPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(givenPass))
	if err != nil {
		return false
	}
	return true
}
func GetPassword(userId int) (string, error) {
	var hashPass string
	qry := "select user_password from user where user_id=?"
	err := db.QueryRow(qry, userId).Scan(&hashPass)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return hashPass, nil
}
func IsNidExist(nid string) (bool, int) {
	var uid int
	query := "SELECT user_id FROM user WHERE nid=?"
	err := db.QueryRow(query, nid).Scan(&uid)
	if err == sql.ErrNoRows {
		return false, 0
	} else if err != nil {
		log.Fatal(err)
		return false, 0
	}
	log.Println("already exist")
	return true, uid
}
func IsNumberExist(num string) (bool, int) {
	var uid int
	query := "SELECT user_id FROM user WHERE mobile_num=?"
	err := db.QueryRow(query, num).Scan(&uid)
	if err == sql.ErrNoRows {
		return false, 0
	} else if err != nil {
		log.Fatal(err)
		return false, 0
	}
	log.Println("already exist")
	return true, uid
}
