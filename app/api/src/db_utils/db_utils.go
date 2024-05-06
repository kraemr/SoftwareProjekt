package db_utils;
import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"src/crypto_utils"
)

const (
	maxOpenConns = 100
	maxIdleConns = 50
	maxLifetime  = 5 * time.Minute
)

var (
	db *sql.DB
)

func InitDB() {
	site_db_pw := os.Getenv("SITE_DB_PASSWORD")
	fmt.Println(site_db_pw)
	cfg := mysql.Config{
		Passwd:               site_db_pw,
		Addr:                 "mariadb:3306",
		Net:                  "tcp",
		DBName:               "SITE_DB",
		AllowNativePasswords: true,
	}
	var err error
	// MYSQL Driver
	fmt.Println(cfg.FormatDSN())
	// DO NOT USE cfg.FormatDSN
	// It literally doesnt work
	db, err = sql.Open("mysql", "root:rootPASSWORD@tcp(mariadb:3306)/SITE_DB?checkConnLiveness=false&maxAllowedPacket=0")
	if err != nil {
		fmt.Println(err)
	}
	// Set connection pool parameters

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(maxLifetime)
}

func testDB() {
	rows, err := db.Query("SELECT * FROM USER")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}


func RegisterUser(email string,password string) bool{
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
	row, err := db.Query("SELECT password from USER where email=? LIMIT 1", email)
	if(err != nil){
		return false
	}
	defer row.Close()
	var hashedPassword string="";
	for row.Next() {
		row.Scan(&hashedPassword)
	}

	fmt.Printf("Hashed PW: %s\n",hashedPassword)
	fmt.Printf("Login PW: %s\n",password)

	correct,err := crypto_utils.CheckPasswordCorrect(password,hashedPassword)
	if(err != nil){
		return	false
	}
	return correct
}