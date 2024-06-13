package attractions;
import(
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"src/sessions"
)


// delete attraction
func delete(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	id := req.URL.Query().Get("id")
	convertedID,err := strconv.ParseInt(id, 10, 64)
	err = RemoveAttraction(convertedID)
	return "{\"success\":true}",err
}

// update existing attraction, check if logged in 
func put(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	var attraction Attraction
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&attraction)
	if(err != nil){
		return "{\"success\":false}",err

	}
	err = UpdateAttraction(attraction)
	if(err != nil){
		return "{\"success\":false}",err

	}
	return "{\"success\":true}",nil
}

// add attraction
// check if logged in
func post(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	var attraction Attraction
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&attraction)
	if(err != nil){
		return "{\"success\":false}",err
	}
	err = InsertAttraction(attraction)
	if(err != nil){
		return "{\"success\":false}",err
	}
	return "{\"success\":true}",nil
}

func get(req *http.Request) (string,error){
	var city string = req.URL.Query().Get("city")
	var title string = req.URL.Query().Get("title")	
	var id string = req.URL.Query().Get("id")
	var category string = req.URL.Query().Get("category")
	var posx string = req.URL.Query().Get("posx")
	var posy string = req.URL.Query().Get("posy")
	var unapproved string = req.URL.Query().Get("unapproved")
	_ = unapproved
	cityIsSet := city != ""
	titleIsSet := title != ""
	idIsSet := id != ""
	categoryIsSet := category != ""
	posxIsSet := posx != ""
	posyIsSet := posy != ""
	unapprovedIsSet := unapproved != ""
	var err error
	var output string
	var attractions []Attraction
	var attraction Attraction

	if(unapprovedIsSet && cityIsSet && sessions.CheckModeratorAccessToCity(req,city)) {
		fmt.Println("filter by unapproved city")
		attractions,err = GetAttractionsUnapprovedCity(city)
	}else if(cityIsSet){	// filter by city
		fmt.Println("filter by city")
		attractions,err = GetAttractionsByCity(city)
	}else if(titleIsSet){ // filter by title 
		attractions,err = GetAttractionsByTitle(title)
	}else if(idIsSet){ // by id
		convertedID := 0
		convertedID,err = strconv.Atoi(id)
		attraction,err = GetAttraction(convertedID)
		if(err != nil){
			output = "{\"info\":\"Attraction with that ID does not exist\"}"
		}
		attractions = append(attractions,attraction)
	}else if(categoryIsSet){ // by category
		attractions,err = GetAttractionsByCategory(category)
	}else if(posxIsSet && posyIsSet){ // by location
		f1, _ := strconv.ParseFloat(posx, 32)
		f2, _ := strconv.ParseFloat(posy, 32)
		attractions,err  = GetAttractionsByPos(float32(f1),float32(f2));
	}else{
		attractions,err = GetAttractions()
	}
	
	if(err != nil){
		return "{\"success\":false,\"info\":\"No Attractions were found\"}",err
	}else{
		json_bytes , json_err := json.Marshal(attractions)	
		if(json_err != nil){
			fmt.Println("error");
		}
		output = string(json_bytes)
		return output,err
	}
}

func HandleAttractionsREST(res http.ResponseWriter, req *http.Request){
	var output string
	var err error
	switch{
		case req.Method == "GET": 
			output,err = get(req)
		case req.Method == "POST":
			output,err = post(req)
		case req.Method == "PUT":
			output,err = put(req)
		case req.Method == "DELETE":
			output,err = delete(req)
	}
	if(err != nil){
		// handle error here, send 500,403,402,401,400 and so on depending on error
		fmt.Fprintf(res,output)
	}else{
		fmt.Fprintf(res, output)
	}
}