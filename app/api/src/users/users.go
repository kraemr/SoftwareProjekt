package users

import (
	"database/sql"
	"fmt"
	"src/db_utils"
)

// User & UserLoginInfo structs - representing the user data just like in the database
type User struct {
	UserId    int64	 `json:userid`
	Email     string `json:email`
	Password  string `json:password`
	City      string `json:city`
	Username  string `json:username`
	Activated string `json:activated`
}

type UserNoPassword struct{
	UserId    int64	 `json:userid`
	Email     string `json:email`
	City      string `json:city`
	Username  string `json:username`
	Activated string `json:activated`
}

type UserLoginInfo struct {
	Email    string `json:email`
	Password string `json:password`
}

// Create ErrNoUser to return when no user is found
var ErrNoUser = fmt.Errorf("No User Found")

func GetUsersByCityAndBanned(city string) ([]User, error) {
	var db *sql.DB = db_utils.DB
	var users []User

	query := "SELECT UserId, Email, Password, City, Username, Activated FROM USER WHERE city=? AND activated='FALSE'"
	rows, err := db.Query(query, city)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var found bool

	// iterate over the rows
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserId, &user.Email, &user.Password, &user.City, &user.Username, &user.Activated)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		found = true
	}

	if !found {
		return nil, ErrNoUser
	}
	return users, nil
}

func GetUserCityById(id int32) (string, error) {
	var db *sql.DB = db_utils.DB
	rows, err := db.Query("SELECT city from USER WHERE id=? LIMIT 1", id)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var city string
	city = ""

	for rows.Next() {
		rows.Scan(&city)
	}

	if city == "" {
		return "", ErrNoUser
	} else {
		return city, nil
	}
}

func GetUserByEmail(email string) (User, error) {
	var db *sql.DB = db_utils.DB
	rows, err := db.Query("SELECT id from USER WHERE email=? LIMIT 1", email)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()
	var id int64
	id = 0

	for rows.Next() {
		rows.Scan(&id)
	}

	if id == 0 {
		return User{}, ErrNoUser
	} else {
		rows, err := db.Query("SELECT id, email, password, city, username from USER WHERE id=? LIMIT 1", id)
		if err != nil {
			return User{}, err
		}
		defer rows.Close()
		var user User
		user.UserId = int64(id)
		for rows.Next() {
			rows.Scan(&user.UserId, &user.Email, &user.Password, &user.City, &user.Username)
		}
		return user, nil
	}
}

func GetUserIdByEmail(email string) (int32, error) {
	var db *sql.DB = db_utils.DB
	rows, err := db.Query("SELECT id from USER WHERE email=? LIMIT 1", email)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var id int32
	id = 0

	for rows.Next() {
		rows.Scan(&id)
	}

	if id == 0 {
		return 0, ErrNoUser
	} else {
		return id, nil
	}
}

// get user by id (assuming function needed)
func GetUserByID(userId int64) (User, error) {
	var db *sql.DB = db_utils.DB
	query := "SELECT id, email, password, city, username FROM USER WHERE id=? LIMIT 1"
	row := db.QueryRow(query, userId)

	var user User
	err := row.Scan(&user.UserId, &user.Email, &user.Password, &user.City, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrNoUser
		}
		return user, err
	}
	return user, nil
}

// delete a user
func DeleteUser(userId int64) error {
	var db *sql.DB = db_utils.DB
	query := "DELETE FROM USER WHERE id=?"
	result, err := db.Exec(query, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNoUser // Make sure ErrNoUser is properly defined in your package
	}

	return nil
}

// update user information
func UpdateUser(newInfo User) error {
	var db *sql.DB = db_utils.DB
	query := "UPDATE USER SET email=?, password=?, city=?, username=? WHERE id=?"
	result, err := db.Exec(query, newInfo.Email, newInfo.Password, newInfo.City, newInfo.Username, newInfo.UserId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNoUser // Make sure ErrNoUser is properly defined in your package
	}

	// Update was successful
	return nil
}

// create a new user
func CreateUser(user User) error {
	var db *sql.DB = db_utils.DB
	query := "INSERT INTO USER (email, password, city, username) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, user.Email, user.Password, user.City, user.Username)
	if err != nil {
		return err
	}
	return nil
}


func GetUsers() ([]UserNoPassword,error){
	var db *sql.DB = db_utils.DB
	rows, err := db.Query("SELECT id,email,city,username,active from USER")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var unpw_list []UserNoPassword
	var unpw UserNoPassword

	no_data := false
	for rows.Next() {
		rows.Scan(&unpw.UserId,&unpw.Email,&unpw.City,&unpw.Username,&unpw.Activated)
		no_data = true
		unpw_list = append(unpw_list,unpw)
	}

	if no_data {
		return nil, ErrNoUser
	} else {
		return unpw_list, nil
	}
}



func GetUsersCity(city string) ([]UserNoPassword,error){
	var db *sql.DB = db_utils.DB
	rows, err := db.Query("SELECT id,email,city,username,active from USER WHERE city=?", city)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var unpw_list []UserNoPassword
	var unpw UserNoPassword

	no_data := true
	for rows.Next() {
		rows.Scan(&unpw.UserId,&unpw.Email,&unpw.City,&unpw.Username,&unpw.Activated)
		no_data = false
		unpw_list = append(unpw_list,unpw)
	}

	if no_data {
		return nil, ErrNoUser
	} else {
		return unpw_list, nil
	}
}