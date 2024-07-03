package apis

import (
	"encoding/json"
	"net/http"
	"src/public_transport"
	"strconv"
)

func routeHandler(fromLat, fromLon, toLat, toLon float64) (string, error) {
	journeys, err := public_transport.FetchFullRouteLongLat(fromLat, fromLon, toLat, toLon)
	if err != nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(journeys)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func HandlePublicTransportREST(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "{\"success\":false, \"message\":\"Invalid request method\"}", http.StatusMethodNotAllowed)
		return
	}

	// Extrahieren der Parameter aus der URL
	params := req.URL.Query()
	fromLatStr, fromLonStr, toLatStr, toLonStr := params.Get("fromLat"), params.Get("fromLon"), params.Get("toLat"), params.Get("toLon")

	// Konvertieren der Parameter in Float64
	fromLat, err1 := strconv.ParseFloat(fromLatStr, 64)
	fromLon, err2 := strconv.ParseFloat(fromLonStr, 64)
	toLat, err3 := strconv.ParseFloat(toLatStr, 64)
	toLon, err4 := strconv.ParseFloat(toLonStr, 64)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		http.Error(res, "{\"success\":false, \"message\":\"Invalid coordinates\"}", http.StatusBadRequest)
		return
	}

	output, err := routeHandler(fromLat, fromLon, toLat, toLon)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Senden der korrekten JSON Antwort
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(output))
}
