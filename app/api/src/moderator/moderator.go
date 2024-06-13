package moderator

import (
	"fmt"
	"src/db_utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"src/users"
	"src/sessions"
)

/*
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
email TEXT,
password TEXT,
city TEXT,
username TEXT
*/

type Moderator struct {
	Id             int32
	Email          string
	Moderates_city string
	Username       string
}

var ErrNoModerator = fmt.Errorf("No Moderators Found")

func GetModeratorById(id int64) (Moderator, error) {
	_ = id
	return Moderator{}, ErrNoModerator
}



func GetModeratorByEmail(email string) (Moderator, error) {
	_ = email
	return Moderator{}, ErrNoModerator
}



func GetModerators(city string) ([]Moderator, error) {
	_ = city
	return nil, ErrNoModerator
}

func UpdateModerator(moderator Moderator) error {
	_ = moderator
	return ErrNoModerator
}

func InsertModerator(moderator Moderator) ([]Moderator, error) {
	_ = moderator
	return nil, ErrNoModerator
}

func DeleteModerator(id int64) error {
	_ = id
	return ErrNoModerator
}

func DisableUser(email string) error{
	var db *sql.DB = db_utils.DB
	query := "UPDATE USER SET active=false WHERE email = ?"
	_,err := db.Exec(query,email)
	if err != nil {
		return err
	}
	return nil
}



type BanInfo struct{
	Email string 
	Reason string  // TODO Save as USER_NOTIFICATION
}



//TODO TEST
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


