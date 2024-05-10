package sessions;
import (
	"src/crypto_utils"
	_	"src/attractions"
	"database/sql"
	"src/db_utils"
)



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

func LoginUser(email string,password string) bool{
// type here is going to be Row instead of Rows
	var db *sql.DB = db_utils.DB
	row, err := db.Query("SELECT password from USER where email=? LIMIT 1", email)
	if(err != nil){
		return false
	}
	defer row.Close()
	var hashedPassword string="";
	for row.Next() {
		row.Scan(&hashedPassword)
	}
	correct,err := crypto_utils.CheckPasswordCorrect(password,hashedPassword)
	if(err != nil){
		return	false
	}
	return correct
}