package users;


import(
	"fmt"
	"src/db_utils"
	"database/sql" 
)

type User struct{
	UserId int64
	Email string
	Password string
	City string
	Username string
	Admin bool
}

type UserLoginInfo struct{
	Email string `json:email`
	Password string `json:password`
}
var ErrNoUser = fmt.Errorf("No User Found")


func GetUserCityById(id int32) (string,error){
	var db *sql.DB = db_utils.DB
	rows, err := db.Query("SELECT city from USER WHERE id=? LIMIT 1", id)
	if(err != nil){
		return "",err
	}
	defer rows.Close()
	var city string
	city = ""
	
	for rows.Next() {
		rows.Scan(&city)
	}

	if(city == ""){
		return "",ErrNoUser
	}else{
		return city,nil
	}
}

func GetUserIdByEmail(email string) (int32,error){
	var db *sql.DB = db_utils.DB
	rows, err := db.Query("SELECT id from USER WHERE email=? LIMIT 1", email)
	if(err != nil){
		return 0,err
	}
	defer rows.Close()
	var id int32 
	id = 0
	
	for rows.Next() {
		rows.Scan(&id)
	}

	if(id==0){
		return 0,ErrNoUser
	}else{
		return id,nil
	}
}