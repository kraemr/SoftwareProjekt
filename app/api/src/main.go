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
	"src/crypto_utils"
	"src/sessions"
//	"io"
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
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Users can register with only email and passwd
// Later on they can add more info if they wish to
func registerUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	_ = decoder
	var user *User_registration = &User_registration{
		Email:"t@g.com",
		Password:"test",
	}
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	prepared_stmt, err := db.Prepare("INSERT INTO USER(email,password) VALUES(?,?)")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	argon2Pw, err := crypto_utils.GetHashedPassword(user.Password);
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user)
	result, err := prepared_stmt.Exec(user.Email, argon2Pw)
	_ = result
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(res, "{\"success\":true}")
}

func loginUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user User_registration
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user)
	
	// type here is going to be Row instead of Rows
    row, err := db.Query("SELECT password from USER where email=? LIMIT 1", user.Email)
	if(err != nil){
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return	
	}
	defer row.Close()
	var hashedPassword string="";
	for row.Next() {
		row.Scan(&hashedPassword)
	}
	
	fmt.Printf("Hashed PW: %s\n",hashedPassword)
	fmt.Printf("Login PW: %s\n",user.Password)
	correct,err := crypto_utils.CheckPasswordCorrect(user.Password,hashedPassword)
	if(err != nil){
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return	
	}
	if(correct == true){
		sessions.StartSession(res,req);
		fmt.Fprintf(res, "{\"success\":true}")
	}else{
		fmt.Fprintf(res, "{\"success\":false}")
	}
}

func findFavoritesForUser(w http.ResponseWriter, r *http.Request) {
}

func findAttractionsNearUser(w http.ResponseWriter, r *http.Request) {
}

func main() {
	argsWithProg := os.Args
	time.Sleep(10 * time.Second) // wait for DB, TODO: make a healthcheck for The DB and in compose wait till healthy
	initDB()
	testDB()
	fmt.Println(argsWithProg)

	// if you want to test outside of the docker then do
	// publicDir := "../../public"
	// in the Docker
	publicDir := "/opt/app/public"
	
	// ########### apis #############
	http.HandleFunc("/api/register", registerUser)
	http.HandleFunc("/api/login", loginUser)
	
	// ########### apis ############

	// start static files server with publicDir as root
	fileServer := http.FileServer(http.Dir(publicDir))
	http.Handle("/", fileServer)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting File server:", err)
	}
	fmt.Println("Http Server is running on port 8000")
}
