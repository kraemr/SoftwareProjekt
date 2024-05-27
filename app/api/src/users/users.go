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
	Admin    bool
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
	// Using GetUserIdByEmail to get the user id
	id, err := GetUserIdByEmail(email)

	// If there is an error, return the error
	if err != nil {
		return User{}, err
	}

	// Get User with id
	return GetUserByID(int64(id)), nil
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
	// TODO: Implement function
	return User{}
}

// send notification to a user
func SendNotification(userId int64, message_json string) error {
	// TODO: Implement function
	return nil
}

// send notification to a city
func SendNotificationToCity(city string, message_json string) error {
	// TODO: Implement function
	return nil
}

// delete a user
func DeleteUser(userId int64) error {
	// TODO: Implement function
	return nil
}

// update user information
func UpdateUser(newInfo User) error {
	// TODO: Implement function
	return nil
}

// create a new user
func CreateUser(user User) error {
	// TODO: Implement function
	return nil
}

// check if credentials are correct
func CheckCorrectCredentials(username string, password string) bool {
	// TODO: Implement function
	return false
}

// get notifications for a user
func GetNotifications(userId int64) []string {
	// TODO: Implement function
	return []string{}
}
