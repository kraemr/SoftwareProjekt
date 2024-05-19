package attractions;
import(
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

func delete(req *http.Request) (string,error){
	return "",nil
}

func put(req *http.Request) (string,error){
	return "",nil

}

func post(req *http.Request) (string,error){
	return "",nil

}

func get(req *http.Request) (string,error){
	var city string = req.URL.Query().Get("city")
	var title string = req.URL.Query().Get("title")	
	var id string = req.URL.Query().Get("id")
	var category string = req.URL.Query().Get("category")
	var posx string = req.URL.Query().Get("posx")
	var posy string = req.URL.Query().Get("posy")
	cityIsSet := city != ""
	titleIsSet := title != ""
	idIsSet := id != ""
	categoryIsSet := category != ""
	posxIsSet := posx != ""
	posyIsSet := posy != ""


	var err error
	var output string
	var attractions []Attraction
	var attraction Attraction

	if(cityIsSet){	// filter by city
		attractions,err = GetAttractionsByCity(city)
	}else if(titleIsSet){ // filter by title 
		attractions,err = GetAttractionsByTitle(title)
	}else if(idIsSet){ // by id
		convertedID := 0
		convertedID,err = strconv.Atoi(id)
		if(err != nil){
			output = "{\"info\":\"Attraction does not exist\"}"
		}
		attraction,err = GetAttraction(convertedID)
		attractions = append(attractions,attraction)
	}else if(categoryIsSet){ // by category
		attractions,err = GetAttractionsByCategory(category)
	}else if(posxIsSet && posyIsSet){ // by location
		f1, ferr := strconv.ParseFloat(posx, 32)
		f2, ferr := strconv.ParseFloat(posy, 32)
		attractions,err  = GetAttractionsByPos(f1,f2);
	}/*else{ // get all with pagination

	}*/

	if(err != nil){

	}else{
		output,err = json.Marshal(attractions)	
		return string(output),err
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
	}else{
		fmt.Fprintf(res, output)
	}
}

/*
// also works
func getAttractionsByCity(res http.ResponseWriter, req *http.Request){
	if(sessions.CheckLoggedIn(req)){
		var city string = req.URL.Query().Get("city")
		attractions,err := attractions.GetAttractionsByCity(city)
		if(err != nil){
			_ = err
			fmt.Fprintf(res, "{\"success\":false}")
		}else{
			_ = attractions
			encoder := json.NewEncoder(res)
			encoder.Encode(attractions)		
		}
	}else{
		// send 403 forbidden, or maybe a redirect to login ?
		fmt.Fprintf(res, "{\"success\":false}")
	}
}


// Works
func getAttractionsByTitle(res http.ResponseWriter, req *http.Request){
	if(sessions.CheckLoggedIn(req)){
		var title string = req.URL.Query().Get("title")
		attractions,err := attractions.GetAttractionsByTitle(title)
		if(err != nil){
			_ = err
			fmt.Fprintf(res, "{\"success\":false}")
		}else{
			_ = attractions
			encoder := json.NewEncoder(res)
			encoder.Encode(attractions)
		}
	}else{
		// send 403 forbidden, or maybe a redirect to login ?
		fmt.Fprintf(res, "{\"success\":false}")
	}
}
*/