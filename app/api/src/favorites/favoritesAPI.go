package favorites;

import(
	"net/http"
	"encoding/json"
	"fmt"
	"src/sessions"
	"strconv"
)


func get(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}

	id := sessions.GetLoggedInUserId(req)
	favorites,err := GetAttractionFavoritesByUserId(int32(id));
	if(err != nil){
		return "{\"success\":false}",err
	}
	json_bytes , json_err := json.Marshal(favorites)	
	if(json_err != nil){
		fmt.Println("error");
		return "{\"success\":false}",json_err
	}
	output := string(json_bytes)
	return output,nil
}


// NOT SUPPORTED
func put(req *http.Request) (string,error){
	_ = req
	return "{\"success\":false,\"info\":\"unsupported Method\"}",nil
}


// Add a Favorite
func post(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	var favorite AttractionFavorite
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&favorite)
	if(err != nil){
		return "{\"success\":false}",err
	}
	err = AddAttractionFavoriteById(favorite.User_id,favorite.Attraction_id)
	if(err != nil){
		return "{\"success\":false}",err
	}
	return "{\"success\":true}",nil
}

func delete(req *http.Request) (string,error){
	if(!sessions.CheckLoggedIn(req)) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}",nil
	}
	id  := req.URL.Query().Get("id")
	convertedID,str_err := strconv.Atoi(id)
	if(str_err != nil){
		return "{\"success\":false}",str_err
	}
	err := DeleteAttractionFavoriteById(int32(convertedID))
	if(err != nil){
		return "{\"success\":false}",err
	}
	return "{\"success\":true}",nil
}


func HandleFavoritesREST(res http.ResponseWriter, req *http.Request){
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