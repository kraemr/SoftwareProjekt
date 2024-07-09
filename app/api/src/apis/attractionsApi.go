package apis

// Attraction REST Api
import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/attractions"
	"src/favorites"
	"src/moderator"
	"src/sessions"
	"strconv"
)

/*
deleteAttraction gets called if a DELETE http-Request is received on /api/attractions
deleteAttraction expects an id inform of a query parameter (http://localhost/api/attractions?id=19191919)
This id identifies the attraction to be deleted
Attractions can only be deleted by a moderator for the corresponding city or by the user who created them
*/
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
/*
This function is called when /api/attractions receives a PUT request
Since this is a REST API the PUT Endpoint is used for updating already existing attractions
This function expects a JSON object in the request body with ALL necessary attributes including Id.
calling it with ?action=approve or disapprove allows moderators to make it visible or invisible to normal users.
*/
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

/*
postAttraction is called when /api/attractions receives a POST request
This creates a new Attraction
This function expects a logged in user sending a JSON Object with all attributes of Attraction struct,
except Stars,Approved,Id,Recommended_count and Added_by as these should not be set by the user
*/
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
	attraction.Stars = 0
	attraction.Approved = false
	attraction.Recommended_count = 0
	attraction.Added_by = sessions.GetLoggedInUserId(req)
	err = attractions.InsertAttraction(attraction)
	if err != nil {
		return "{\"success\":false}", err
	}
	return "{\"success\":true}", nil
}


/*
getAttraction is called when requesting Attraction data
This function is called when /api/attractions receives a GET Request
getAttraction allows to filter by a bunch of attributes like city,title,id,category,approval and so on
unapproved attractions should only be returned when a moderator sends a request
When the Attractions are retrieved from the Database the FavoriteCount is inserted into the attraction struct
*/
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
		fmt.Println(unapproved);
		attraction_list, err = attractions.GetAttractionsUnapproved()
	} else {
		attraction_list, err = attractions.GetAttractions()
	}

	// not very pretty but gets the job done, might slow down at about 1000
	for i:=0;i<len(attraction_list);i+=1 {
		attraction_list[i].Recommended_count,err = favorites.GetAttractionFavoriteCountByAttractionId( attraction_list[i].Id )
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



/*
This is the Callback function for /api/attractions
Depending on Request Method different functions are called
Every function returns JSON-String and an error value that is nil on success
*/
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
