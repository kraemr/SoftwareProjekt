package main
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"src/sessions"
	"src/db_utils"
)

type User_registration struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Users can register with only email and passwd
// Later on they can add more info if they wish to
func registerUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	_ = decoder
	var user *User_registration = &User_registration{
		Email:"t@g.com",
		Password:"test",
	}
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	success := db_utils.RegisterUser(user.Email,user.Password);
	if(success == true){
		fmt.Fprintf(res, "{\"success\":true}")
	}else{
		fmt.Fprintf(res, "{\"success\":false}")
	}
}

func loginUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user User_registration
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user)
	
	correct := db_utils.LoginUser(user.Email,user.Password);
	if(correct == true){
		sessions.StartSession(res,req);
		fmt.Fprintf(res, "{\"success\":true}")
	}else{
		fmt.Fprintf(res, "{\"success\":false}")
	}
}

func findFavoritesForUser(w http.ResponseWriter, r *http.Request) {
}

func findAttractionsNearUser(w http.ResponseWriter, r *http.Request) {
}

func main() {
	argsWithProg := os.Args
	time.Sleep(10 * time.Second) // wait for DB, TODO: make a healthcheck for The DB and in compose wait till healthy
	db_utils.InitDB()
	fmt.Println(argsWithProg)
	// if you want to test outside of the docker then do
	// publicDir := "../../public"
	// in the Docker
	publicDir := "/opt/app/public"
	
	// ########### apis #############
	http.HandleFunc("/api/register", registerUser)
	http.HandleFunc("/api/login", loginUser)
	
	// ########### apis ############

	// start static files server with publicDir as root
	fileServer := http.FileServer(http.Dir(publicDir))
	http.Handle("/", fileServer)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting File server:", err)
	}
	fmt.Println("Http Server is running on port 8000")
}
