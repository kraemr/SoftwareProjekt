package users;
import (
	"net/http"
)

func get(){

}

func post(){

}

func delete(){

}

func put(){

}



/*
// Users can register with only email and passwd
// Later on they can add more info if they wish to, that would be with a PUT api call

func registerUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	_ = decoder
	var user *users.UserLoginInfo = &users.UserLoginInfo{
		Email:"",
		Password:"",
	}
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	success := sessions.RegisterUser(user.Email,user.Password);
	if(success == true){
		fmt.Fprintf(res, "{\"success\":true}")
	}else{
		fmt.Fprintf(res, "{\"success\":false}")
	}
}
*/

func HandleUsersREST(res http.ResponseWriter, req *http.Request){
}