package moderator
import (
	"fmt"
	"src/db_utils"
	"database/sql"
)
/*
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
email TEXT,
password TEXT,
city TEXT,
username TEXT
*/
type Moderator struct {
	Id             int32
	Email          string
	City 		   string
	Username       string
}
var ErrNoModerator = fmt.Errorf("No Moderators Found")

func GetModeratorById(id int64) (Moderator, error) {
	var db *sql.DB = db_utils.DB
	sql := "Select * from CITY_MODERATOR where id = ?"
	row, err := db.Query(sql,id)
	if(err != nil){
		return Moderator{},err
	}
	defer row.Close()
	var m Moderator
	nodata_found := true
	for row.Next() {
		var pw string
		row.Scan(&m.Id,&m.Email,&pw,&m.City,&m.Username)
		_ = pw // pw is not used
		nodata_found = false

	}
	if(nodata_found){
		return Moderator{}, ErrNoModerator
	}else{
		return m,nil;
	}
}

func GetModeratorByEmail(email string) (Moderator, error) {
	var db *sql.DB = db_utils.DB
	sql := "Select * from CITY_MODERATOR where email = ?"
	row, err := db.Query(sql,email)
	if(err != nil){
		return Moderator{},err
	}
	defer row.Close()
	var m Moderator
	nodata_found := true
	for row.Next() {
		var pw string
		row.Scan(&m.Id,&m.Email,&pw,&m.City,&m.Username)
		_ = pw // pw is not used
		nodata_found = false

	}
	if(nodata_found){
		return Moderator{}, ErrNoModerator
	}else{
		return m,nil;
	}
}



func GetModeratorsCity(city string) ([]Moderator, error) {
	var db *sql.DB = db_utils.DB
	sql := "Select * from CITY_MODERATOR where city = ?"
	row, err := db.Query(sql,city)
	if(err != nil){
		return nil,err
	}
	defer row.Close()
	var m Moderator
	var mods []Moderator
	nodata_found := true
	for row.Next() {
		var pw string
		row.Scan(&m.Id,&m.Email,&pw,&m.City,&m.Username)
		_ = pw // pw is not used
		mods = append(mods,m)
		nodata_found = false
	}
	if(nodata_found){
		return nil, ErrNoModerator
	}else{
		return mods,nil;
	}
}

func UpdateModerator(moderator Moderator) error {
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare(" UPDATE CITY_MODERATOR SET id=?,email=?,city=?,username=? WHERE id=?")
	if(err != nil){
		fmt.Println("Couldnt Insert Moderator")
		return err
	}
	result,err := prepared_stmt.Exec(
		moderator.Id,
		moderator.Email,
		moderator.City,
		moderator.Username,
		moderator.Id)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

func InsertModerator(moderator Moderator) (error) {
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("INSERT INTO CITY_MODERATOR(id,email,city,username) VALUES(?,?,?,?)")
	if(err != nil){
		fmt.Println("Couldnt Insert Moderator")
		return err
	}
	result,err := prepared_stmt.Exec(
		moderator.Id,
		moderator.Email,
		moderator.City,
		moderator.Username)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

func DeleteModerator(id int64) error {
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare(" Delete FROM CITY_MODERATOR WHERE id=?")
	if(err != nil){
		fmt.Println("Couldnt Insert Moderator")
		return err
	}
	result,err := prepared_stmt.Exec(id)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

func DisableUser(email string) error{
	var db *sql.DB = db_utils.DB
	query := "UPDATE USER SET active=false WHERE email = ?"
	_,err := db.Exec(query,email)
	if err != nil {
		return err
	}
	return nil
}