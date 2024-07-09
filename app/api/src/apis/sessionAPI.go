package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/moderator"
	"src/sessions"
	"src/users"
)

func CheckUserLoggedIn(res http.ResponseWriter, req *http.Request) {
	if sessions.CheckLoggedIn(req) == true {
		fmt.Fprintf(res, "{\"success\":true}")
	} else {
		fmt.Fprintf(res, "{\"success\":false}")
	}
}



/*
This function receives the Data from the moderatorlogin and checks 
that the Information is correct
if it is a session is created where logged_in = true
and certain variables like moderator moderates_city are set
*/
func LoginModerator(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var modInfo users.UserLoginInfo
	err := decoder.Decode(&modInfo)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(modInfo)

	correct := sessions.LoginModerator(modInfo.Email, modInfo.Password)
	if correct == true {
		// get Id
		mod, uerr := moderator.GetModeratorByEmail(modInfo.Email)
		if uerr != nil {
			fmt.Println("loginUser: couldnt get user by id")
			return
		}
		sessions.StartModeratorSession(res, req, mod.Id, mod.City)
		fmt.Fprintf(res, "{\"success\":true}")
	} else {
		fmt.Fprintf(res, "{\"success\":false}")
	}
}

/*
This function receives the Data from the User Login and checks 
that the Information is correct
if it is a session is created where logged_in = true
and certain variables like userId are set for later use
*/
func LoginUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user users.UserLoginInfo
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user)

	correct := sessions.LoginUser(user.Email, user.Password)
	if correct == true {
		// get Id
		id, uerr := users.GetUserIdByEmail(user.Email)
		if uerr != nil {
			fmt.Println("loginUser: couldnt get user by id")
			return
		}
		sessions.StartSession(res, req, id)
		fmt.Fprintf(res, "{\"success\":true}")
	} else {
		fmt.Fprintf(res, "{\"success\":false}")
	}
}


/*NOT IMPLEMENTED*/
func LogoutAPI(res http.ResponseWriter, req *http.Request){
	sessions.Logout(req)
	fmt.Fprintf(res,"\"success\":true");
}
