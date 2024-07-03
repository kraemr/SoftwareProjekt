package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/sessions"
	"src/moderator"
	"strconv"
	"src/attractions"
)

// deleteAttraction attraction
func deleteAttraction(req *http.Request) (string,error){
	id := req.URL.Query().Get("id")
	convertedID,err := strconv.ParseInt(id, 10, 64)
	a1,err1 := attractions.GetAttraction(convertedID)
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
			err = attractions.RemoveAttraction(convertedID)
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

	err = attractions.RemoveAttraction(convertedID)
	return "{\"success\":true}", err
}




// is not really how you would normally do a PUT api
func putModeratingAction(req *http.Request) error{
	var action string = req.URL.Query().Get("action")
	var id string = req.URL.Query().Get("id")
	convertedID,err := strconv.ParseInt(id, 10, 64)
	_ = err
	switch action {
		case "approve":
			attractions.ChangeAttractionApproval(true,convertedID);
		case "disapprove":
			attractions.ChangeAttractionApproval(false,convertedID);
	}
	return nil

}



// update existing attraction, check if logged in and added_by id is the users
func putAttraction(req *http.Request) (string,error){
	if(req.URL.Query().Get("action") != ""){
		e := putModeratingAction(req);
		if(e != nil){
			return "{\"success\":false}",e
		}else{
			return "{\"success\":true}",nil
		}
	}
	
	
	var attraction attractions.Attraction
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
		
		a,err1 := attractions.GetAttraction(attraction.Id)
		if(err1 != nil){
			return "{\"success\":false}",err 
		}

		if(mod.City != a.City){
			return "{\"success\":false}",nil 
		}

		err = attractions.UpdateAttraction(attraction)
		if(err != nil){
			return "{\"success\":false}",err
		}
	}

	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	userid := sessions.GetLoggedInUserId(req)
	a1,err1 := attractions.GetAttraction(attraction.Id)
	if(err1 != nil){
		return "{\"success\":false}",err1
	}

	if(a1.Added_by != userid){
		return "{\"success\":false}",nil
	}

	if( err == nil ){
		err = attractions.UpdateAttraction(attraction)
		if(err != nil){
			return "{\"success\":false}",err
		}
		return "{\"success\":true}",nil
	}
	return "{\"success\":false}",nil

}

// add attraction
// check if logged in, attraction will not be approved, moderator must look at it first
func postAttraction(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	var attraction attractions.Attraction
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&attraction)
	if err != nil {
		return "{\"success\":false}", err
	}
	err = attractions.InsertAttraction(attraction)
	if err != nil {
		return "{\"success\":false}", err
	}
	return "{\"success\":true}", nil
}

func getAttraction(req *http.Request) (string, error) {
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
	var attraction_list []attractions.Attraction
	var attraction attractions.Attraction

	
	if cityIsSet && categoryIsSet{
		attraction_list, err = attractions.GetAttractionsByCityAndType(city,category)
	} else if unapprovedIsSet && cityIsSet && sessions.CheckModeratorAccessToCity(req, city) {
		fmt.Println("filter by unapproved city")
		attraction_list, err = attractions.GetAttractionsUnapprovedCity(city)
	} else if cityIsSet { // filter by city
		fmt.Println("filter by city")
		attraction_list, err = attractions.GetAttractionsByCity(city)
	} else if titleIsSet { // filter by title
		attraction_list, err = attractions.GetAttractionsByTitle(title)
	} else if idIsSet { // by id
		convertedID := int64(0)
		convertedID, err = strconv.ParseInt(id, 10, 64)
		attraction, err = attractions.GetAttraction(convertedID)
		if err != nil {
			output = "{\"info\":\"Attraction with that ID does not exist\"}"
		}
		attraction_list = append(attraction_list, attraction)
	} else if categoryIsSet { // by category
		attraction_list, err = attractions.GetAttractionsByCategory(category)
	} else if posxIsSet && posyIsSet { // by location
		f1, _ := strconv.ParseFloat(posx, 32)
		f2, _ := strconv.ParseFloat(posy, 32)
		attraction_list, err = attractions.GetAttractionsByPos(float32(f1), float32(f2))
	} else if(unapprovedIsSet){
		attraction_list, err = attractions.GetAttractionsUnapproved()
	} else {
		attraction_list, err = attractions.GetAttractions()
	}

	if err != nil {
		return "{\"success\":false,\"info\":\"No Attractions were found\"}", err
	} else {
		json_bytes, json_err := json.Marshal(attraction_list)
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
		output, err = getAttraction(req)
	case req.Method == "POST":
		output, err = postAttraction(req)
	case req.Method == "PUT":
		output, err = putAttraction(req)
	case req.Method == "DELETE":
		output, err = deleteAttraction(req)
	}
	if err != nil {
		// handle error here, send 500,403,402,401,400 and so on depending on error
		fmt.Fprintf(res, output)
	} else {
		fmt.Fprintf(res, output)
	}
}
