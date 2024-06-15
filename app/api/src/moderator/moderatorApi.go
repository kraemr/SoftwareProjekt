package moderator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"src/sessions"
)


// moderators can only delete themselves
// Site Admin can delete Moderator
func delete(req *http.Request) (string, error) {
	return "{\"success\":false}", nil 
	id := req.URL.Query().Get("id")
	convertedID, err := strconv.ParseInt(id, 10, 64)
	err = DeleteModerator(convertedID)
	return "{\"success\":true}", err
}


// moderators can update ONLY himself
// Site admins can update Everything he wants
func put(req *http.Request) (string, error) {
	if(!sessions.CheckLoggedIn(req)){
		return "{\"success\":false}", nil
	}
	var moderator Moderator
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&moderator)
	if err != nil {
		return "{\"success\":false}", err
	}
	err = UpdateModerator(moderator)
	if err != nil {
		return "{\"success\":false}", err

	}
	return "{\"success\":true}", nil
}



// ONLY Site Admins can add new Moderators
func post(req *http.Request) (string, error) {
	var moderator Moderator
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&moderator)
	if err != nil {
		return "{\"success\":false}", err
	}
	err = InsertModerator(moderator)
	if err != nil {
		return "{\"success\":false}", err
	}
	return "{\"success\":true}", nil
}

// Only Moderators can get Info about Moderators
func get(req *http.Request) (string, error) {
	var city string = req.URL.Query().Get("city")
	var id string = req.URL.Query().Get("id")
	var email string = req.URL.Query().Get("email")

	id_is_set := id != ""
	city_is_set := city != ""
	email_is_set := email != ""
	var mods []Moderator
	var mod Moderator
	var output string
	var err error

	if id_is_set {
		var convertedID int64
		convertedID, err = strconv.ParseInt(id, 10, 64)
		if(err != nil){
			return "{\"success\":false}", err
		}
		mod,err = GetModeratorById(convertedID)
		if(err != nil){
			return "{\"success\":false}", err
		}
		mods = append(mods,mod)
	} else if city_is_set {
		mods,err = GetModeratorsCity(city)
		if(err != nil){
			return "{\"success\":false}", err
		}
	}else if email_is_set {
		mod,err = GetModeratorByEmail(email)
		if(err != nil){
			return "{\"success\":false}", err
		}
		mods = append(mods,mod)
	}


	json_bytes , json_err := json.Marshal(mods)	
	if(json_err != nil){
		fmt.Println("error");
	}
	output = string(json_bytes)
	return output,nil
}

func HandleModeratorsREST(res http.ResponseWriter, req *http.Request) {
	var output string
	var err error
	switch {
	case req.Method == "GET":
		output, err = get(req)
	case req.Method == "POST":
		output, err = post(req)
	case req.Method == "PUT":
		output, err = put(req)
	case req.Method == "DELETE":
		output, err = delete(req)
	}
	if err != nil {
		// handle error here, send 500,403,402,401,400 and so on depending on error
		fmt.Fprintf(res, output)
	} else {
		fmt.Fprintf(res, output)
	}
}
