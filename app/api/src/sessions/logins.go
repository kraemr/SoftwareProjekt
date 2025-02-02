package sessions

import (
	"database/sql"
	"fmt"
	"src/crypto_utils"
	"src/db_utils"
)

/*
Registers a User with email and saves PasswordHash
*/
func RegisterUser(email string,password string) bool{
	var db *sql.DB = db_utils.DB
	prepared_stmt, err := db.Prepare("INSERT INTO USER(email,password) VALUES(?,?)")
	if err != nil {
		return false
	}
	argon2Pw, err := crypto_utils.GetHashedPassword(password);
	if err != nil {
		return false 
	}
	result, err := prepared_stmt.Exec(email, argon2Pw)
	_ = result
	if err != nil {
		return false
	}
	return true
}

/*
This function checks if the user supplied password hashed is equal to the one in the database
*/
func LoginUser(email string,password string) bool{
// type here is going to be Row instead of Rows
	var db *sql.DB = db_utils.DB
	row, err := db.Query("SELECT password from USER where email=? LIMIT 1", email)
	if(err != nil){
		return false
	}
	defer row.Close()
	var hashedPassword string
	for row.Next() {
		row.Scan(&hashedPassword)
	}
	correct,err := crypto_utils.CheckPasswordCorrect(password,hashedPassword)
	if(err != nil){
		return	false
	}
	return correct
}


/*
This function checks if the moderator supplied password hashed is equal to the one in the database
*/
func LoginModerator(email string,password string) bool{
	// type here is going to be Row instead of Rows
		var db *sql.DB = db_utils.DB
		row, err := db.Query("SELECT password from CITY_MODERATOR where email=? LIMIT 1", email)
		if(err != nil){
			fmt.Println("Email for mod doesnt exist")
			return false
		}
		defer row.Close()
		var hashedPassword string
		for row.Next() {
			row.Scan(&hashedPassword)
		}
		correct,err := crypto_utils.CheckPasswordCorrect(password,hashedPassword)
		if(err != nil){
			return	false
		}
		return correct
	}