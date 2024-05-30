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

func StartModeratorSession(w http.ResponseWriter, r *http.Request,id int32,city string){
	session, err := store.Get(r, "sessionid")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["logged_in"] = true
	session.Values["id"] = id
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
		return false
	}
	// test if the type is correct here 
	if logged_in, ok := session.Values["logged_in"].(bool); !ok {
		fmt.Println("session was nil or unexpected type")
		_ = logged_in
		return false
	}
	
	if logged_in, ok := session.Values["logged_in"].(bool); ok{
		return logged_in
	}
	return false
}


func CheckModeratorAccessToCity(r *http.Request	,city string) bool{
	session, err := store.Get(r, "sessionid")
	if(err != nil){
		return false
	}

	if logged_in, ok := session.Values["logged_in"].(bool); !ok {
		fmt.Println("session was nil or unexpected type")
		return false
	}

	if city , ok := session.Values["moderator_city"].(string); !ok{
		fmt.Println("session was nil or unexpected type")
		return false
	}
	return true;
}
