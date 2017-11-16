package main

import (
	"database/sql"
	"log"
	"skeletor/utils"
)

func saveUserProfile(profile *Profile) {
	profile.Password = utils.HashPassword(profile.Password)
	err := session.QueryRow(`INSERT INTO profile ( 
		firstname, 
		lastname, 
		username, 
		email, 
		title, 
		password, 
		mobilenumber
	) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		profile.Firstname,
		profile.Lastname,
		profile.Username,
		profile.Email,
		profile.Title,
		profile.Password,
		profile.MobileNumber).Scan(&profile.Id)
	if err != nil {
		log.Print(err)
	}
	profile.Password = ""
}

func queryUserCredential(profile *Profile) bool {
	result := false

	err := session.QueryRow(`SELECT 
		firstname, 
		lastname, 
		username, 
		email, 
		title, 
		mobilenumber FROM profile WHERE 
		username = $1 AND password = $2`, profile.Username, profile.Password).Scan(&profile.Firstname,
		&profile.Lastname,
		&profile.Username,
		&profile.Email,
		&profile.Title,
		&profile.MobileNumber)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		result = true
	}

	return result
}
