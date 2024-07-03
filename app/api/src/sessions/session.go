package sessions;
import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"fmt"
)
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func GetLoggedInUserId( r *http.Request) int32{
	session, err := store.Get(r, "sessionid")
	if(err != nil){
		return -1
	}
	id,ok := session.Values["id"].(int32)
	if(!ok){
		return -1
	}
	return id
}

func UserIsBanned(r *http.Request) bool{
	session, err := store.Get(r, "sessionid")
	if err != nil {
		return false
	}	
	if banned, ok := session.Values["banned"].(bool); ok{
		return banned
	}else{
		return false
	}
}


func StartSession(w http.ResponseWriter, r *http.Request,id int32){
	session, err := store.Get(r, "sessionid")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["logged_in"] = true
	session.Values["id"] = id
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func StartModeratorSession(w http.ResponseWriter, r *http.Request,moderator_id int32,city string){
	session, err := store.Get(r, "sessionid")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["logged_in"] = true
	session.Values["id"] = moderator_id
	session.Values["moderator_city"]  = city
	session.Values["moderator"] = true

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CheckLoggedIn(r *http.Request) bool {
	session, err := store.Get(r, "sessionid")
	if(err != nil){
		fmt.Println("ssession doesnt exist")
		return false
	}
	// test if the type is correct here 
	if logged_in, ok := session.Values["logged_in"].(bool); !ok {
		fmt.Println("session was nil or unexpected type")
		_ = logged_in
		return false
	}
	
	if logged_in, ok := session.Values["logged_in"].(bool); ok{
		fmt.Printf("logged_in was %b\n",logged_in);
		return logged_in
	}
	return false
}



func CheckModeratorLoggedIn(r *http.Request) bool {
	session, err := store.Get(r, "sessionid")
	if(err != nil){
		return false
	}

	if logged_in, ok := session.Values["logged_in"].(bool); !ok {
		fmt.Println("session was nil or unexpected type")
		_ = logged_in
		return false
	}
	logged_in_state := false
	
	if logged_in, ok := session.Values["logged_in"].(bool); ok{
		logged_in_state = logged_in
	}

	if is_moderator, ok := session.Values["moderator"].(bool); ok{
		return is_moderator && logged_in_state
	}

	return false
}



func CheckModeratorAccessToCity(r *http.Request	, _city string) bool{
	session, err := store.Get(r, "sessionid")
	if(err != nil){
		return false
	}

	if logged_in, ok := session.Values["logged_in"].(bool); !ok {
		_ = logged_in
		fmt.Println("session was nil or unexpected type")
		return false
	}

	if city , ok := session.Values["moderator_city"].(string); !ok {
		_ = city
		fmt.Println("session was nil or unexpected type")
		return false
	}else if(city == _city){
		return true
	}
	return false
}


func Logout(r *http.Request){
	session, err := store.Get(r, "sessionid")
	_ = err
	_ = session
	session.Values["logged_in"] = false
	session.Values["moderator"] = false
	session.Values["id"] = 0
	session.Values["moderator_city"] = ""


	session.Options.MaxAge=-1
}
