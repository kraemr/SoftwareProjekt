package db_utils;
import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/go-sql-driver/mysql"
	"database/sql"
)

const (
	maxOpenConns = 100
	maxIdleConns = 50
	maxLifetime  = 5 * time.Minute
)

var (
	DB *sql.DB
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
	DB, err = sql.Open("mysql", "root:rootPASSWORD@tcp(mariadb:3306)/SITE_DB?checkConnLiveness=false&maxAllowedPacket=0")
	if err != nil {
		fmt.Println(err)
	}
	// Set connection pool parameters

	DB.SetMaxOpenConns(maxOpenConns)
	DB.SetMaxIdleConns(maxIdleConns)
	DB.SetConnMaxLifetime(maxLifetime)
}

func testDB() {
	rows, err := DB.Query("SELECT * FROM USER")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

func GetCurrentDate() string{
	now := time.Now()
	return now.Format("2006-01-02")
}

