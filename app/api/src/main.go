package main
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	_ "time"
	"src/sessions"
	"src/db_utils"
	"src/attractions"
	"src/users"
	"src/notifications"
)


// Users can register with only email and passwd
// Later on they can add more info if they wish to
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

func loginUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user users.UserLoginInfo
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user)
	
	correct := sessions.LoginUser(user.Email,user.Password);
	if(correct == true){
		sessions.StartSession(res,req);
		fmt.Fprintf(res, "{\"success\":true}")
	}else{
		fmt.Fprintf(res, "{\"success\":false}")
	}
}

func startServer(port string){
	fmt.Println("Http Server is running on port: " + port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		fmt.Println("Error starting File server:", err)
	}
}


func main() {
	argsWithProg := os.Args
	db_utils.InitDB()
	fmt.Println(argsWithProg)
	if(len(argsWithProg ) > 1 && argsWithProg[1] == "test"){
		// run tests
	}
	publicDir := "/opt/app/public"
	// ########### apis #############
	http.HandleFunc("/api/register", registerUser)
	http.HandleFunc("/api/login", loginUser)
	http.HandleFunc("/api/attractions",attractions.HandleAttractionsREST)
	http.HandleFunc("/api/users",users.HandleUsersREST)

	// ########### apis ############
	// start static files server with publicDir as root
	fileServer := http.FileServer(http.Dir(publicDir))
	http.Handle("/", fileServer)
	go notifications.StartNotificationServer("8080","/notifications")
	startServer("8000") // keeps running i.e blocks execution
}
