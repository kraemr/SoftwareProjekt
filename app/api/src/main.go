package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	maxOpenConns = 100
	maxIdleConns = 50
	maxLifetime  = 5 * time.Minute
)

var (
	db *sql.DB
)

func initDB() {
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

type User_registration struct {
	email    string
	password string
}

// Users can register with only email and passwd
// Later on they can add more info if they wish to
func registerUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user User_registration
	err := decoder.Decode(&user)
	if err != nil {
	}

	// TODO: HASH PASSWORD FIRST
	prepared_stmt, err := db.Prepare("INSERT INTO USER(email,password) VALUES(?,?);")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := prepared_stmt.Exec(user.email, user.password)
	_ = result
	fmt.Fprintf(res, "{\"success\":true}")
}

func loginUser(w http.ResponseWriter, r *http.Request) {
}

func findFavoritesForUser(w http.ResponseWriter, r *http.Request) {
}

func findAttractionsNearUser(w http.ResponseWriter, r *http.Request) {
}

func main() {
	argsWithProg := os.Args

	time.Sleep(10 * time.Second)

	initDB()
	testDB()
	fmt.Println(argsWithProg)
	// if you want to test outside of the docker then do
	// publicDir := "../../public"
	// in the Docker
	publicDir := "/opt/app/public"
	http.HandleFunc("/api/register", registerUser)
	fileServer := http.FileServer(http.Dir(publicDir))
	http.Handle("/", fileServer)
	fmt.Println("Server is running on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting File server:", err)
	}
}
