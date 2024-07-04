package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/sessions"
	"src/users"
	"strconv"
)

// delete user
func delete(req *http.Request) (string, error) {
	id := req.URL.Query().Get("id")
	convertedID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", err
	}
	err = users.DeleteUser(convertedID)
	return "{\"success\":true}", err
}

// update existing user
func put(req *http.Request) (string, error) {
	var user users.User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		return "", err
	}
	err = users.UpdateUser(user)
	return "{\"success\":true}", err
}

// Add user
func post(req *http.Request) (string, error) {
	var user users.User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		return "", err
	}
	err = users.CreateUser(user)
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
	var user users.User

	if idIsSet { // by id
		convertedID, err := strconv.ParseInt(inID, 10, 64)
		if err != nil {
			return "{\"info\":\"User does not exist\"}", err
		}
		user, err = users.GetUserByID(convertedID)
		if err != nil {
			return "{\"success\":false,\"info\":\"Error getting user by ID\"}", err
		}
		outputBytes, err := json.Marshal(user)
		if err != nil {
			return "{\"success\":false,\"info\":\"Error marshalling user data\"}", err
		}
		output = string(outputBytes)
	} else if emailIsSet { // by email
		user, err = users.GetUserByEmail(inEmail)
		if err != nil {
			return "{\"success\":false,\"info\":\"Error marshalling user data\"}", err
		}
		outputBytes, err := json.Marshal(user)
		if err != nil {
			return "{\"success\":false,\"info\":\"Error marshalling user data\"}", err
		}
		output = string(outputBytes)
	} else { // By current logged in user (session)
		id := sessions.GetLoggedInUserId(req)
		user, err = users.GetUserByID(int64(id))
		if err != nil {
			return "{\"success\":false,\"info\":\"Error marshalling user data\"}", err
		}
		outputBytes, err := json.Marshal(user)
		if err != nil {
			return "{\"success\":false,\"info\":\"Error marshalling user data\"}", err
		}
		output = string(outputBytes)
	}
	return output, err
}

// Rest API Handler for users
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
