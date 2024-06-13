package recommendations;

import (
	"net/http"
	"fmt"
	"encoding/json"
	"src/attractions"
	"src/sessions"
)

func get(req *http.Request) (string,error){
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

func delete(req *http.Request) (string,error){
	return "{\"success\":false,\"info\":\"unsupported Method\"}",nil
}

func put(req *http.Request) (string,error){
	return "{\"success\":false,\"info\":\"unsupported Method\"}",nil
}

func post(req *http.Request) (string,error){
	return "{\"success\":false,\"info\":\"unsupported Method\"}",nil
}

func HandleRecommendationsREST(res http.ResponseWriter, req *http.Request){
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