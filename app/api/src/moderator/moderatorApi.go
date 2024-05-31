package moderator;

package attractions;
import(
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)


// delete attraction
func delete(req *http.Request) (string,error){
	id := req.URL.Query().Get("id")
	convertedID,err := strconv.ParseInt(id, 10, 64)
	err = DeleteModerator(convertedID)
	return "{\"success\":true}",err
}

// update existing attraction, check if logged in 
func put(req *http.Request) (string,error){
	var moderator Moderator
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&moderator)
	if(err != nil){
		return "{\"success\":false}",err
	}
	err = UpdateModerator(moderator)
	if(err != nil){
		return "{\"success\":false}",err

	}
	return "{\"success\":true}",nil
}

// add attraction
// check if logged in
func post(req *http.Request) (string,error){
	var moderator Moderator
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&moderator)
	if(err != nil){
		return "{\"success\":false}",err
	}
	err = InsertModerator(moderator)
	if(err != nil){
		return "{\"success\":false}",err
	}
	return "{\"success\":true}",nil
}

func get(req *http.Request) (string,error){
	var city string = req.URL.Query().Get("city")
	var id string = req.URL.Query().Get("id")
	id_is_set := id != ""
	city_is_set := city != ""
	if(id_is_set){

	}else if(city_is_set){

	}
	return "",err
}

func HandleModeratorsREST(res http.ResponseWriter, req *http.Request){
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