package attractions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/sessions"
	"src/moderator"
)

// delete attraction
func delete(req *http.Request) (string,error){
	id := req.URL.Query().Get("id")
	convertedID,err := strconv.ParseInt(id, 10, 64)
	a1,err1 := GetAttraction(int(convertedID))
	if(err1 != nil){
		return "{\"success\":false}",nil
	}

	if(err != nil){
		return "{\"success\":false}",err
	}

	if(sessions.CheckModeratorLoggedIn(req)){
		mod,err := moderator.GetModeratorById(convertedID)
		if(err != nil){
			return "{\"success\":false,\"info\":\"Not Logged in\"}",err 
		}

		if(a1.City == mod.City){
			err = RemoveAttraction(convertedID)
			return "{\"success\":true}",nil
		}
	}

	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	userid := sessions.GetLoggedInUserId(req)
	
	if(a1.Added_by != userid){
		return "{\"success\":false}",nil
	}

	err = RemoveAttraction(convertedID)
	return "{\"success\":true}", err
}

// update existing attraction, check if logged in and added_by id is the users
func put(req *http.Request) (string,error){
	var attraction Attraction
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&attraction)
	if(err != nil){
		return "{\"success\":false}",err
	}

	if(sessions.CheckModeratorLoggedIn(req)){
		id := sessions.GetLoggedInUserId(req)
		
		mod,err := moderator.GetModeratorById(int64(id))
		if(err != nil){
			return "{\"success\":false,\"info\":\"Not Logged in\"}",err 
		}
		
		a,err1 := GetAttraction(int(attraction.Id))
		if(err1 != nil){
			return "{\"success\":false}",err 
		}

		if(mod.City != a.City){
			return "{\"success\":false}",nil 
		}

		err = UpdateAttraction(attraction)
		if(err != nil){
			return "{\"success\":false}",err
		}
	}
<<<<<<< HEAD

	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	userid := sessions.GetLoggedInUserId(req)
	a1,err1 := GetAttraction(int(attraction.Id))
	if(err1 != nil){
		return "{\"success\":false}",err1
	}

	if(a1.Added_by != userid){
		return "{\"success\":false}",nil
	}

	if( err == nil ){
		err = UpdateAttraction(attraction)
		if(err != nil){
			return "{\"success\":false}",err
		}
		return "{\"success\":true}",nil
	}
	return "{\"success\":false}",nil

}

// add attraction
// check if logged in, attraction will not be approved, moderator must look at it first
func post(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
=======
	return "{\"success\":true}", nil
}

// add attraction
// check if logged in
func post(req *http.Request) (string, error) {
	if !sessions.CheckLoggedIn(req) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}", nil
>>>>>>> 6243376aa92208686ec8c6b720b1b82763f261b4
	}
	var attraction Attraction
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&attraction)
	if err != nil {
		return "{\"success\":false}", err
	}
	err = InsertAttraction(attraction)
	if err != nil {
		return "{\"success\":false}", err
	}
	return "{\"success\":true}", nil
}

func get(req *http.Request) (string, error) {
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

	if unapprovedIsSet && cityIsSet && sessions.CheckModeratorAccessToCity(req, city) {
		fmt.Println("filter by unapproved city")
		attractions, err = GetAttractionsUnapprovedCity(city)
	} else if cityIsSet { // filter by city
		fmt.Println("filter by city")
		attractions, err = GetAttractionsByCity(city)
	} else if titleIsSet { // filter by title
		attractions, err = GetAttractionsByTitle(title)
	} else if idIsSet { // by id
		convertedID := int64(0)
		convertedID, err = strconv.ParseInt(id, 10, 64)
		attraction, err = GetAttraction(convertedID)
		if err != nil {
			output = "{\"info\":\"Attraction with that ID does not exist\"}"
		}
		attractions = append(attractions, attraction)
	} else if categoryIsSet { // by category
		attractions, err = GetAttractionsByCategory(category)
	} else if posxIsSet && posyIsSet { // by location
		f1, _ := strconv.ParseFloat(posx, 32)
		f2, _ := strconv.ParseFloat(posy, 32)
		attractions, err = GetAttractionsByPos(float32(f1), float32(f2))
	} else {
		attractions, err = GetAttractions()
	}

	if err != nil {
		return "{\"success\":false,\"info\":\"No Attractions were found\"}", err
	} else {
		json_bytes, json_err := json.Marshal(attractions)
		if json_err != nil {
			fmt.Println("error")
		}
		output = string(json_bytes)
		return output, err
	}
}

func HandleAttractionsREST(res http.ResponseWriter, req *http.Request) {
	var output string
	var err error
	switch {
	case req.Method == "GET":
		output, err = get(req)
	case req.Method == "POST":
		output, err = post(req)
	case req.Method == "PUT":
		output, err = put(req)
	case req.Method == "DELETE":
		output, err = delete(req)
	}
	if err != nil {
		// handle error here, send 500,403,402,401,400 and so on depending on error
		fmt.Fprintf(res, output)
	} else {
		fmt.Fprintf(res, output)
	}
}
