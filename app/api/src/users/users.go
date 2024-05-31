package users

import (
	"database/sql"
	"fmt"
	"src/db_utils"
)

type User struct {
	UserId   int64
	Email    string
	Password string
	City     string
	Username string
}

type UserLoginInfo struct {
	Email    string `json:email`
	Password string `json:password`
}

var ErrNoUser = fmt.Errorf("No User Found")

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
func GetUserByID(userId int64) User {
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
    _, err := db.Exec(query, userId)
    if err != nil {
        return err
    }
    return nil
}

// update user information
func UpdateUser(newInfo User) error {
	var db *sql.DB = db_utils.db
	query := "UPDATE USER SET email=?, password=?, city=?, username=? WHERE id=?"
    _, err := db.Exec(query, newInfo.Email, newInfo.Password, newInfo.City, newInfo.Username, newInfo.UserId)
    if err != nil {
        return err
    }
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