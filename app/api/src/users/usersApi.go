package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// delete user
func delete(req *http.Request) (string, error) {
	id := req.URL.Query().Get("id")
	convertedID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", err
	}
	err = DeleteUser(convertedID)
	return "{\"success\":true}", err
}

// update existing user
func put(req *http.Request) (string, error) {
	var user User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		return "", err
	}
	err = UpdateUser(user)
	return "{\"success\":true}", err
}

// Add user
func post(req *http.Request) (string, error) {
	var user User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		return "", err
	}
	err = CreateUser(user)
	return "{\"success\":true}", err
}

// get existing user by id or email
func get(req *http.Request) (string, error) {
	var inID string = req.URL.Query().Get("id")
	var inEmail string = req.URL.Query().Get("email")

	idIsSet := inID != ""
	emailIsSet := inEmail != ""

	var err error
	var output string
	var user User

	if idIsSet { // by id
		convertedID, err := strconv.ParseInt(inID, 10, 64)
		if err != nil {
			return "{\"info\":\"User does not exist\"}", err
		}
		user = GetUserByID(convertedID)
		outputBytes, err := json.Marshal(user)
		if err != nil {
			return "{\"success\":false,\"info\":\"Error marshalling user data\"}", err
		}
		output = string(outputBytes)
	} else if emailIsSet { // by email
		user, err = GetUserByEmail(inEmail)
		if err != nil {
			return "{\"success\":false,\"info\":\"Error marshalling user data\"}", err
		}
		outputBytes, err := json.Marshal(user)
		if err != nil {
			return "{\"success\":false,\"info\":\"Error marshalling user data\"}", err
		}
		output = string(outputBytes)
	} else {
		return "{\"success\":false,\"info\":\"No user id provided\"}", nil
	}
	return output, err
}

func HandleUsersREST(res http.ResponseWriter, req *http.Request) {
	var output string
	var err error
	switch req.Method {
	case "GET":
		output, err = get(req)
	case "POST":
		output, err = post(req)
	case "PUT":
		output, err = put(req)
	case "DELETE":
		output, err = delete(req)
	}
	if err != nil {
		// handle error here, send 500,403,402,401,400 and so on depending on error
		fmt.Fprintf(res, "{\"success\":false,\"info\":\"%v\"}", err)
	} else {
		fmt.Fprintf(res, output)
	}
}
