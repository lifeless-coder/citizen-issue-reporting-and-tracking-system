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
func ComparePass(hashpass, givenPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(givenPass))
	if err != nil {
		return false
	}
	return true
}
func IsAdminMailExist(mail string) (bool, int) {
	var uid int
	query := "SELECT admin_id FROM admin WHERE email=?"
	err := db.QueryRow(query, mail).Scan(&uid)
	if err == sql.ErrNoRows {
		return false, 0 //na thakle false
	} else if err != nil {
		log.Fatal(err)
		return false, 0
	}
	log.Println("already exist")
	return true, uid

}
func GetAdminPassword(userId int) (string, error) {
	var hashPass string
	qry := "select admin_password from admin where admin_id=?"
	err := db.QueryRow(qry, userId).Scan(&hashPass)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return hashPass, nil
}
func InsertIntoAdmin(uname, mail, pass string, addedby int) {
	query := "INSERT INTO admin" +
		"(username, email, admin_password, added_by)" +
		"VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, uname, mail, pass, addedby)
	if err != nil {
		log.Println("Error inserting user: ", err)
	}
}
