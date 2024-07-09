package moderator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/sessions"
	"src/users"
);

type BanInfo struct{
	Email string 
	Reason string  // TODO Save as USER_NOTIFICATION
}

/*
Ban a user by his email and with a given Reason
This sets the users active boolean to false
*/
func BanUser(res http.ResponseWriter, req *http.Request){
	if(req.Method == "PUT"){
		var ban BanInfo
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&ban)
		if(err != nil){
			fmt.Fprintf(res,"{\"success\":false,\"info\":\"invalid Data\"}")
			return;
		}else{
			errUser := DisableUser(ban.Email);
			if(errUser != nil){
				fmt.Fprintf(res,"{\"success\":false,\"info\":\"User does not exist\"}")
				return;
			}
			fmt.Fprintf(res,"{\"success\":true}")
		}
	}
}


/*
Returns List of BannedUsers for given City 
If the current session is a Moderator session and the moderator has access to the city
*/
func GetBannedUsers(res http.ResponseWriter, req *http.Request){
	var city string = req.URL.Query().Get("city")
	
	if(	!sessions.CheckModeratorAccessToCity(req,city) ){
		fmt.Fprintf(res,"{\"success\":false}")
		return
	}

	if(req.Method == "GET"){
		user_list,err := users.GetUsersByCityAndBanned(city)
		if err != nil {
			fmt.Fprintf(res,"{\"success\":false}")
			return
		}else{
			json_bytes , json_err := json.Marshal(user_list)
			if(json_err != nil){
				fmt.Println("json error")
			}
			output := string(json_bytes)
			fmt.Fprintf(res,output)
		}
	}
}


