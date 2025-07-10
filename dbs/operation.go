package dbs

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InsertIntoUser(uname, nid, mobile, pass, gender, prof, bday, pic, mail string) {
	query := "INSERT INTO user" +
		"(username, nid, mobile_num, user_password, gender, profession, birthdate, picture, email)" +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, uname, nid, mobile, pass, gender, prof, bday, pic, mail)
	if err != nil {
		log.Println("Error inserting user: ", err)
	}
}
