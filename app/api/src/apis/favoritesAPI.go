package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/favorites"
	"src/sessions"
	"strconv"
)

func getFavorite(req *http.Request) (string, error) {
	if !sessions.CheckLoggedIn(req) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}", nil
	}

	id := sessions.GetLoggedInUserId(req)
	favorite_list, err := favorites.GetAttractionFavoritesByUserId(int64(id))
	if err != nil {
		return "{\"success\":false}", err
	}
	json_bytes, json_err := json.Marshal(favorite_list)
	if json_err != nil {
		fmt.Println("error")
		return "{\"success\":false}", json_err
	}
	output := string(json_bytes)
	return output, nil
}

// NOT SUPPORTED
func putFavorite(req *http.Request) (string, error) {
	_ = req
	return "{\"success\":false,\"info\":\"unsupported Method\"}", nil
}

// Add a Favorite
func postFavorite(req *http.Request) (string, error) {
	if !sessions.CheckLoggedIn(req) {
		return "{\"success\":false,\"info\":\"Not Logged in\"}", nil
	}
	var favorite favorites.AttractionFavorite
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&favorite)
	if err != nil {
		return "{\"success\":false}", err
	}
	err = favorites.AddAttractionFavoriteById(favorite.User_id, favorite.Attraction_id)
	if err != nil {
		return "{\"success\":false}", err
	}
	return "{\"success\":true}", nil
}

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
