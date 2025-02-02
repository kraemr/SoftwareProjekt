package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/attractions"
	"src/sessions"
)

/*
This function gets called when /apu/recommendations receives a GET request
the user can specify which city and category he wants recommendations for
In an actual Production environment we woul have to have a pretty complex recommendation algo
to give users what they want.
Currently this just returns the best rated attractions in that city/category
*/

func getRecommendations(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}

	var city string = req.URL.Query().Get("city")
	var category string = req.URL.Query().Get("category")
	
	cityIsSet := city != ""
	categoryIsSet := category != ""
	
	id := sessions.GetLoggedInUserId(req)
	if(cityIsSet && categoryIsSet){
		recs,err := attractions.GetRecommendationForUser(id,city,category)
		if(err != nil){
			return "{\"success\":false}",err
		}

		json_bytes , json_err := json.Marshal(recs)	
		if(json_err != nil){
			fmt.Println("error");
			return "{\"success\":false}",json_err
		}
		output := string(json_bytes)
		return output,nil
	}else{
		return "{\"success\":false}",nil
	}
}

func deleteRecommendations(req *http.Request) (string,error){
	_ =req
	return "{\"success\":false,\"info\":\"unsupported Method\"}",nil
}

func putRecommendations(req *http.Request) (string,error){
	_ =req
	return "{\"success\":false,\"info\":\"unsupported Method\"}",nil
}

func postRecommendations(req *http.Request) (string,error){
	_ =req
	return "{\"success\":false,\"info\":\"unsupported Method\"}",nil
}



/*
This is the Callback function for /api/recommendations
Depending on Request Method different functions are called
Every function returns JSON-String and an error value that is nil on success
*/
func HandleRecommendationsREST(res http.ResponseWriter, req *http.Request){
	var output string
	var err error
	switch{
		case req.Method == "GET": 
			output,err = getRecommendations(req)
		case req.Method == "POST":
			output,err = postRecommendations(req)
		case req.Method == "PUT":
			output,err = putRecommendations(req)
		case req.Method == "DELETE":
			output,err = deleteRecommendations(req)
	}
	if(err != nil){
		// handle error here, send 500,403,402,401,400 and so on depending on error
		fmt.Fprintf(res,output)
	}else{
		fmt.Fprintf(res, output)
	}
}