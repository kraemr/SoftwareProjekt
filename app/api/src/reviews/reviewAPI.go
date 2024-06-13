package reviews;
import(
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)


func delete(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	id := req.URL.Query().Get("id")
	convertedID,err := strconv.ParseInt(id, 10, 64)
	if(err != nil){
		return "{\"success\":false}",err
	}
	err = DeleteReview(convertedID)
	if(err != nil){
		return "{\"success\":false}",err
	}else{
		return "{\"success\":true}",nil
	}
}

func put(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	var review Review
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&review)
	if(err != nil){
		return "{\"success\":false}",err

	}
	err = UpdateReview(review)
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
	var review Review
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&review)
	if(err != nil){
		return "{\"success\":false}",err

	}
	err = InsertReview(review)
	if(err != nil){
		return "{\"success\":false}",err

	}
	return "{\"success\":true}",nil
}

func get(req *http.Request) (string,error){
	var user_id string = req.URL.Query().Get("user_id")
	var attraction_id string = req.URL.Query().Get("attraction_id")
	user_id_set := user_id != ""
	attraction_id_set := attraction_id != ""

	var reviews []Review
	var err error
	

	if(user_id_set){
		convertedID := 0
		convertedID,err = strconv.Atoi(user_id)
		reviews,err = GetReviewsByUserId(int32(convertedID))
		if(err != nil){
			return "{\"success\":false}",nil
		}
	}else if(attraction_id_set){
		convertedID := 0
		convertedID,err = strconv.Atoi(attraction_id)
		reviews,err = GetReviewsByAttractionId(int32( convertedID))
		if(err != nil){
			fmt.Println("failed GetReviewsByAttractionId");
			return "{\"success\":false}",nil
		}
	}

	json_bytes , json_err := json.Marshal(reviews)	
	if(json_err != nil){
		fmt.Println("error");
		return  "{\"success\":false}",nil
	}
	output := string(json_bytes)
	return output,err
}

func HandleReviewREST(res http.ResponseWriter, req *http.Request){
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