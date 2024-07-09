package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/favorites"
	"src/sessions"
	"strconv"
)

/*
getFavorite gets called if /api/favorites receives a GET Request
getFavorite returns Loggedin Users Favorites when no parameters are present
if action parameter == count and attraction_id exists then this returns the FavoriteCount for the attraction
*/
func getFavorite(req *http.Request) (string, error) {	
	if(req.URL.Query().Get("action") == "count" && req.URL.Query().Get("attraction_id") != ""){
		a_id,e := strconv.ParseInt( req.URL.Query().Get("attraction_id"), 10, 64);
		if(e != nil){
			return "{\"success\":false,\"info\":\"attraction_id missing\"}", e
		}
		count,err := favorites.GetAttractionFavoriteCountByAttractionId(a_id);
		if(err != nil){
			return "{\"success\":false,\"info\":\"Couldnt Count Favorites\"}", err
		}
		return fmt.Sprintf("{\"favorite_count\":%d}", count) , nil
	}
	

	if !sessions.CheckLoggedIn(req) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}", nil
	}
	
	id := sessions.GetLoggedInUserId(req)
	favorite_list, err := favorites.GetAttractionFavoritesByUserId(int64(id))
	if err != nil {
		return "{\"success\":false,\"info\":\"GetAttractionFavoritesByUserId failed\"}", err
	}
	json_bytes, json_err := json.Marshal(favorite_list)
	if json_err != nil {
		fmt.Println("error")
		return "{\"success\":false,\"info\":\"json decode failed\"}", json_err
	}
	output := string(json_bytes)
	return output, nil
}

// NOT SUPPORTED
func putFavorite(req *http.Request) (string, error) {
	_ = req
	return "{\"success\":false,\"info\":\"unsupported Method\"}", nil
}


/*
This is not really conformant with REST principles, but it was way easier to NOT manage this Logic in the frontend
postFavorite is used to set an Attraction to a Favorite of the loggedin user or to unset it
if the favorite exists, then it deletes the favorite
if the favorite doesnt exist, then it creates a new favorite
*/
func postFavorite(req *http.Request) (string, error) {
	if !sessions.CheckLoggedIn(req) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}", nil
	}
	var favorite favorites.AttractionFavorite
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&favorite)
	if err != nil {
		return "{\"success\":false,\"info\":\"favorite decode fail\"}", err
	}
	fmt.Println(favorite)
	b,e := favorites.CheckFavoriteExists(favorite.Attraction_id, favorite.User_id)
	if(b == true && e == nil){
		delete_err := favorites.DeleteAttractionFavoriteByAttractionId(favorite.Attraction_id,favorite.User_id);
		if(delete_err != nil){
			return "{\"success\":false,\"info\":\"did not delete favorite\"}",nil;
		}
		return "{\"success\":true,\"info\":\"deleted favorite\"}",nil;
	}

	err = favorites.AddAttractionFavoriteById(favorite.User_id, favorite.Attraction_id)
	if err != nil {
		fmt.Println(err)
		return "{\"success\":false,\"info\":\"AddAttractionFavoriteById failed\"}", err
	}
	return "{\"success\":true,\"info\":\"added favorite\"}", nil
}

/*
This function gets called when /api/favorites receives a DELETE
This function expects an id that identifies a favorite
if it exists and the user is logged in AND it is actually his, then it is deleted
*/
func deleteFavorite(req *http.Request) (string, error) {
	if !sessions.CheckLoggedIn(req) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}", nil
	}
	id := req.URL.Query().Get("id")
	convertedID, str_err := strconv.Atoi(id)
	if str_err != nil {
		return "{\"success\":false}", str_err
	}
	err := favorites.DeleteAttractionFavoriteById(int64(convertedID))
	if err != nil {
		return "{\"success\":false}", err
	}
	return "{\"success\":true}", nil
}

/*
This is the Callback function for /api/favorites
Depending on Request Method different functions are called
Every function returns JSON-String and an error value that is nil on success
*/
func HandleFavoritesREST(res http.ResponseWriter, req *http.Request) {
	var output string
	var err error
	switch {
	case req.Method == "GET":
		output, err = getFavorite(req)
	case req.Method == "POST":
		output, err = postFavorite(req)
	case req.Method == "PUT":
		output, err = putFavorite(req)
	case req.Method == "DELETE":
		output, err = deleteFavorite(req)
	}
	if err != nil {
		// handle error here, send 500,403,402,401,400 and so on depending on error
		fmt.Fprintf(res, output)
	} else {
		fmt.Fprintf(res, output)
	}
}
