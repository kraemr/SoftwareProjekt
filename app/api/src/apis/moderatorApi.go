package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/moderator"
	"src/sessions"
	"strconv"
)

// moderators can only delete themselves
// Site Admin can deleteFavorite Moderator
// No proper checks implemented for now ... SO just return success:false
func deleteModerator(req *http.Request) (string, error) {
	return "{\"success\":false}", nil 
	id := req.URL.Query().Get("id")
	convertedID, err := strconv.ParseInt(id, 10, 64)
	err = moderator.DeleteModerator(convertedID)
	return "{\"success\":true}", err
}


// moderators can update ONLY himself
// Site admins can update Everything he wants
func putModerator(req *http.Request) (string, error) {
	if(!sessions.CheckLoggedIn(req)){
		return "{\"success\":false}", nil
	}
	var mod moderator.Moderator
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&mod)
	if err != nil {
		return "{\"success\":false}", err
	}
	err = moderator.UpdateModerator(mod)
	if err != nil {
		return "{\"success\":false}", err

	}
	return "{\"success\":true}", nil
}



// ONLY Site Admins can add new Moderators
func postModerator(req *http.Request) (string, error) {
	var mod moderator.Moderator
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&mod)
	if err != nil {
		return "{\"success\":false}", err
	}
	err = moderator.InsertModerator(mod)
	if err != nil {
		return "{\"success\":false}", err
	}
	return "{\"success\":true}", nil
}

// Only Moderators can getModerator Info about Moderators
func getModerator(req *http.Request) (string, error) {
	var city string = req.URL.Query().Get("city")
	var id string = req.URL.Query().Get("id")
	var email string = req.URL.Query().Get("email")

	id_is_set := id != ""
	city_is_set := city != ""
	email_is_set := email != ""
	var mods []moderator.Moderator
	var mod moderator.Moderator
	var output string
	var err error

	if id_is_set {
		var convertedID int64
		convertedID, err = strconv.ParseInt(id, 10, 64)
		if(err != nil){
			return "{\"success\":false}", err
		}
		mod,err = moderator.GetModeratorById(convertedID)
		if(err != nil){
			return "{\"success\":false}", err
		}
		mods = append(mods,mod)
	} else if city_is_set {
		mods,err = moderator.GetModeratorsCity(city)
		if(err != nil){
			return "{\"success\":false}", err
		}
	}else if email_is_set {
		mod,err = moderator.GetModeratorByEmail(email)
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
		output, err = getModerator(req)
	case req.Method == "POST":
		output, err = postModerator(req)
	case req.Method == "PUT":
		output, err = putModerator(req)
	case req.Method == "DELETE":
		output, err = deleteModerator(req)
	}
	if err != nil {
		// handle error here, send 500,403,402,401,400 and so on depending on error
		fmt.Fprintf(res, output)
	} else {
		fmt.Fprintf(res, output)
	}
}
