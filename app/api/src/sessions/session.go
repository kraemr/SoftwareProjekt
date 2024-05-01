package sessions;
import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"fmt"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func StartSession(w http.ResponseWriter, r *http.Request){
	session, err := store.Get(r, "sessionid")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["logged_in"] = true
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