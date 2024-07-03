package apis;
import (
	"src/moderator"
	"net/http"
	"encoding/json"
	"src/sessions"
	"fmt"
	"src/users"
)

func CheckUserLoggedIn(res http.ResponseWriter, req *http.Request) {
	if sessions.CheckLoggedIn(req) == true {
		fmt.Fprintf(res, "{\"success\":true}")
	} else {
		fmt.Fprintf(res, "{\"success\":false}")
	}
}

func LoginModerator(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var modInfo users.UserLoginInfo
	err := decoder.Decode(&modInfo)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(modInfo)

	correct := sessions.LoginUser(modInfo.Email, modInfo.Password)
	if correct == true {
		// get Id
		mod, uerr := moderator.GetModeratorByEmail(modInfo.Email)
		if uerr != nil {
			fmt.Println("loginUser: couldnt get user by id")
			return
		}
		sessions.StartModeratorSession(res, req, mod.Id, mod.Email)
		fmt.Fprintf(res, "{\"success\":true}")
	} else {
		fmt.Fprintf(res, "{\"success\":false}")
	}
}

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
