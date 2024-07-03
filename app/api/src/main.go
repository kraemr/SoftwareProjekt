package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"src/db_utils"
	"src/moderator"
	"src/notifications"
	"src/public_transport"
	"src/apis"
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

var categories [7]string

func handleCategories(res http.ResponseWriter, req *http.Request) {
	json_bytes, json_err := json.Marshal(categories)
	if json_err != nil {
		fmt.Fprintf(res, "{\"success\":false}")
	}
	output := string(json_bytes)
	fmt.Fprintf(res, output)
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

	http.HandleFunc("/api/login", apis.LoginUser)
	http.HandleFunc("/api/logged_in", apis.CheckUserLoggedIn)
	http.HandleFunc("/api/login_moderator", apis.LoginModerator)

	http.HandleFunc("/api/ban", moderator.BanUser)
	http.HandleFunc("/api/banned", moderator.GetBannedUsers)
	http.HandleFunc("/api/categories", handleCategories)

	// ########### Rest apis #############
	http.HandleFunc("/api/attractions", apis.HandleAttractionsREST)
	http.HandleFunc("/api/users", apis.HandleUsersREST)
	http.HandleFunc("/api/favorites", apis.HandleFavoritesREST)
	http.HandleFunc("/api/recommendations", apis.HandleRecommendationsREST)
	http.HandleFunc("/api/reviews", apis.HandleReviewREST)
	http.HandleFunc("/api/moderators", apis.HandleModeratorsREST)
	// ########### Rest apis ###########

	// start static files server with publicDir as root
	fileServer := http.FileServer(http.Dir(publicDir))
	http.Handle("/", fileServer)
	go notifications.StartNotificationServer("8080", "/notifications")

	//_,_ = recommendations.GetRecommendationForUser(100,"Berlin","Landmark");

	testPublicTransport()
	startServer("8000") // keeps running i.e blocks execution
}
