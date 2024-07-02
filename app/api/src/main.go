package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"src/attractions"
	"src/db_utils"
	"src/favorites"
	"src/moderator"
	"src/notifications"
	"src/public_transport"
	"src/recommendations"
	"src/reviews"
	"src/sessions"
	"src/users"
	_ "time"
)

func testPublicTransport() {
	// Definiert die Koordinaten für den Startpunkt und den Zielpunkt.
	fromLat, fromLon := 49.9179102, 8.3430285 // Beispielkoordinaten für irgendwo in Nackenheim
	toLat, toLon := 49.987809, 8.2272517      // Beispielkoordinaten für Lucy-Hillebrand-Straße, Mainz
	// Holt die beste Route zwischen den beiden Standorten.
	journeys, err := public_transport.FetchFullRouteLongLat(fromLat, fromLon, toLat, toLon)

	if err != nil {
		fmt.Println("Error fetching route:", err)
		return
	}
	// Gibt die gefundene Route aus.
	for _, journey := range journeys {
		fmt.Println(journey)
	}
}

func checkUserLoggedIn(res http.ResponseWriter, req *http.Request) {
	if sessions.CheckLoggedIn(req) == true {
		fmt.Fprintf(res, "{\"success\":true}")
	} else {
		fmt.Fprintf(res, "{\"success\":false}")
	}
}

func loginUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user users.UserLoginInfo
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user)

	correct := sessions.LoginUser(user.Email, user.Password)
	if correct == true {
		// get Id
		id, uerr := users.GetUserIdByEmail(user.Email)
		if uerr != nil {
			fmt.Println("loginUser: couldnt get user by id")
			return
		}
		sessions.StartSession(res, req, id)
		fmt.Fprintf(res, "{\"success\":true}")
	} else {
		fmt.Fprintf(res, "{\"success\":false}")
	}
}

var categories [7]string

func handleCategories(res http.ResponseWriter, req *http.Request) {
	json_bytes, json_err := json.Marshal(categories)
	if json_err != nil {
		fmt.Fprintf(res, "{\"success\":false}")
	}
	output := string(json_bytes)
	fmt.Fprintf(res, output)
}

func loginModerator(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var modInfo users.UserLoginInfo
	err := decoder.Decode(&modInfo)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(modInfo)

	correct := sessions.LoginUser(modInfo.Email, modInfo.Password)
	if correct == true {
		// get Id
		mod, uerr := moderator.GetModeratorByEmail(modInfo.Email)
		if uerr != nil {
			fmt.Println("loginUser: couldnt get user by id")
			return
		}
		sessions.StartModeratorSession(res, req, mod.Id, mod.Email)
		fmt.Fprintf(res, "{\"success\":true}")
	} else {
		fmt.Fprintf(res, "{\"success\":false}")
	}
}

func startServer(port string) {
	fmt.Println("Http Server is running on port: " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting File server:", err)
	}
}

func main() {
	argsWithProg := os.Args
	db_utils.InitDB()
	fmt.Println(argsWithProg)
	if len(argsWithProg) > 1 && argsWithProg[1] == "test" {
		// run tests
	}
	publicDir := "/opt/app/public"
	categories = [...]string{"Monument", "Castle", "Cathedral", "Palace", "Museum", "Mountain", "Park"}

	http.HandleFunc("/api/login", loginUser)
	http.HandleFunc("/api/logged_in", checkUserLoggedIn)
	http.HandleFunc("/api/login_moderator", loginModerator)

	http.HandleFunc("/api/ban", moderator.BanUser)
	http.HandleFunc("/api/banned", moderator.GetBannedUsers)
	http.HandleFunc("/api/categories", handleCategories)

	// ########### Rest apis #############
	http.HandleFunc("/api/attractions", attractions.HandleAttractionsREST)
	http.HandleFunc("/api/users", users.HandleUsersREST)
	http.HandleFunc("/api/favorites", favorites.HandleFavoritesREST)
	http.HandleFunc("/api/recommendations", recommendations.HandleRecommendationsREST)
	http.HandleFunc("/api/reviews", reviews.HandleReviewREST)
	http.HandleFunc("/api/moderators", moderator.HandleModeratorsREST)
	// ########### Rest apis ###########

	// start static files server with publicDir as root
	fileServer := http.FileServer(http.Dir(publicDir))
	http.Handle("/", fileServer)
	go notifications.StartNotificationServer("8080", "/notifications")

	//_,_ = recommendations.GetRecommendationForUser(100,"Berlin","Landmark");

	testPublicTransport()
	startServer("8000") // keeps running i.e blocks execution
}
