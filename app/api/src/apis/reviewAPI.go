package apis

import(
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"src/sessions"
	"src/reviews"
)


func deleteReview(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	id := req.URL.Query().Get("id")
	convertedID,err := strconv.ParseInt(id, 10, 64)
	if(err != nil){
		return "{\"success\":false}",err
	}
	err = reviews.DeleteReview(convertedID)
	if(err != nil){
		return "{\"success\":false}",err
	}else{
		return "{\"success\":true}",nil
	}
}

func putReview(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	var review reviews.Review
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&review)
	if(err != nil){
		return "{\"success\":false}",err
	}
	err = reviews.UpdateReview(review)
	if(err != nil){
		return "{\"success\":false}",err

	}
	return "{\"success\":true}",nil
}

// add attraction
// check if logged in
func postReview(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	var review reviews.Review
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&review)
	if(err != nil){
		return "{\"success\":false}",err

	}
	err = reviews.InsertReview(review)
	if(err != nil){
		return "{\"success\":false}",err

	}
	return "{\"success\":true}",nil
}

func getReview(req *http.Request) (string,error){
	var user_id string = req.URL.Query().Get("user_id")
	var attraction_id string = req.URL.Query().Get("attraction_id")
	user_id_set := user_id != ""
	attraction_id_set := attraction_id != ""
	var reviews_list []reviews.Review
	var err error

	if(req.URL.Query().Get("action") != "" && attraction_id_set){
		convertedID := 0
		convertedID,err = strconv.Atoi(attraction_id)
		if(err != nil){
			return "{\"success\":false}",err
		}
		stars,e :=  reviews.GetStarsForAttraction(int32(convertedID))
		if(e != nil){
			return "{\"success\":false,\"info\":\"no reviews\"}",err
		}
		stars_string := strconv.FormatFloat(float64(stars), 'f', -1, 32)
		return "{\"stars\":" +  stars_string + "}",nil ;
	}

	if(user_id_set){
		convertedID := 0
		convertedID,err = strconv.Atoi(user_id)
		reviews_list,err = reviews.GetReviewsByUserId(int32(convertedID))
		if(err != nil){
			return "{\"success\":false}",nil
		}
	}else if(attraction_id_set){
		convertedID := 0
		convertedID,err = strconv.Atoi(attraction_id)
		fmt.Println(convertedID)
		reviews_list,err = reviews.GetReviewsByAttractionId(int32( convertedID))
		if(err != nil){
			fmt.Println("failed GetReviewsByAttractionId");
			fmt.Println(err)
			return "{\"success\":false}",nil
		}
	}

	json_bytes , json_err := json.Marshal(reviews_list)	
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
			output,err = getReview(req)
		case req.Method == "POST":
			output,err = postReview(req)
		case req.Method == "PUT":
			output,err = putReview(req)
		case req.Method == "DELETE":
			output,err = deleteReview(req)
	}
	if(err != nil){
		// handle error here, send 500,403,402,401,400 and so on depending on error
		fmt.Fprintf(res,output)
	}else{
		fmt.Fprintf(res, output)
	}
}