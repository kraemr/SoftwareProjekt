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
)


// Users can register with only email and passwd
// Later on they can add more info if they wish to
func registerUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	_ = decoder
	var user *users.UserLoginInfo = &users.UserLoginInfo{
		Email:"t@g.com",
		Password:"test",
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




func findFavoritesForUser(w http.ResponseWriter, r *http.Request) {
}

func findAttractionsNearUser(w http.ResponseWriter, r *http.Request) {
}

/*
func test(){
	test_attraction := attractions.Attraction{}
	test_attraction.Title = "testTitle"
	test_attraction.Type = "testType"
	test_attraction.Info = "testInfo"
	test_attraction.Recommended_count = 100000
	test_attraction.PosX = 20.0
	test_attraction.PosY = 8.542
	test_attraction.City = "Oppenheim"
	err := attractions.InsertAttraction(test_attraction)
	if(err != nil){
		fmt.Println(err.Error())
	}
	err = attractions.ChangeAttractionApproval(true,1)
	if(err != nil){
		fmt.Println(err.Error())
	}else{
		fmt.Println("Approvedd Attraction")
	}

	attr,err2 := attractions.GetAttraction(1)
	if(err2 != nil){
		fmt.Println(err2.Error())
	}else{
		fmt.Println(attr)
	}


	attraction_list,err3 := attractions.GetAttractionsByCity("Oppenheim")
	if(err3 != nil){
		fmt.Printf("%v\n",err3.Error())
	}else{
		fmt.Printf("attraction_list for Oppenheim: %d\n",len(attraction_list))
	}

	attraction_list2,err4 := attractions.GetAttractionsByTitle("est")
	if(err3 != nil){
		fmt.Printf("%v\n",err4.Error())
	}else{
		fmt.Printf("attraction_list for Oppenheim: %d\n",len(attraction_list2))
	}

}
*/

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
