package public_transport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://v6.db.transport.rest"

// Location repräsentiert einen Standort mit seinen Details.
type Location struct {
	Type     string `json:"type"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
}

// Journey repräsentiert eine Reise mit mehreren Abschnitten (Legs).
type Journey struct {
	Type string `json:"type"`
	Legs []struct {
		TripID    string `json:"tripId"`
		Direction string `json:"direction"`
		Line      struct {
			Type    string `json:"type"`
			ID      string `json:"id"`
			Name    string `json:"name"`
			Mode    string `json:"mode"`
			Product string `json:"product"`
		} `json:"line"`
		Origin struct {
			Type     string `json:"type"`
			ID       string `json:"id"`
			Name     string `json:"name"`
			Location struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"location"`
		} `json:"origin"`
		Departure        string `json:"departure"`
		PlannedDeparture string `json:"plannedDeparture"`
		Destination      struct {
			Type     string `json:"type"`
			ID       string `json:"id"`
			Name     string `json:"name"`
			Location struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"location"`
		} `json:"destination"`
		Arrival        string `json:"arrival"`
		PlannedArrival string `json:"plannedArrival"`
	} `json:"legs"`
}

// fetchLocationID ruft die Standort-ID für gegebene Koordinaten ab.
func fetchLocationID(latitude, longitude float64) (string, error) {
	var locationID string
	resp, err := http.Get(fmt.Sprintf("%s/locations/nearby?latitude=%f&longitude=%f&results=1", baseURL, latitude, longitude))
	if err != nil {
		return locationID, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return locationID, err
	}
	var locations []Location
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return locationID, err
	}
	if len(locations) > 0 {
		locationID = locations[0].ID
	}
	return locationID, nil
}

// fetchJourneys ruft die besten Reisen zwischen zwei Standorten ab.
func fetchJourneys(fromID, toID string) ([]Journey, error) {
	var journeys []Journey
	departureTime := time.Now().Add(24 * time.Hour).Format(time.RFC3339)
	resp, err := http.Get(fmt.Sprintf("%s/journeys?from=%s&to=%s&departure=%s&results=1", baseURL, fromID, toID, url.QueryEscape(departureTime)))
	if err != nil {
		return journeys, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return journeys, err
	}
	var data struct {
		Journeys []Journey `json:"journeys"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return journeys, err
	}
	journeys = data.Journeys
	return journeys, nil
}

func FetchFullRouteLongLat(fromLat, fromLon, toLat, toLon float64) ([]Journey, error) {
	// Holt die Standort-ID für den Zielpunkt.
	fromID, err := fetchLocationID(fromLat, fromLon)
	if err != nil {
		err = fmt.Errorf("Error fetching to location ID: %w", err)
		return nil, err
	}

	// Holt die Standort-ID für den Zielpunkt.
	toID, err := fetchLocationID(toLat, toLon)
	if err != nil {
		err = fmt.Errorf("Error fetching to location ID: %w", err)
		return nil, err
	}

	// Holt die besten Reisen von der Start-ID zur Ziel-ID.
	journeys, err := fetchJourneys(fromID, toID)
	if err != nil {
		err = fmt.Errorf("Error fetching journeys: %w", err)
		return nil, err
	}

	if len(journeys) == 0 {
		fmt.Println("No journeys found")
		err = fmt.Errorf("No journeys found")
		return nil, err
	}

	// Gefundenen Reisen zurückgeben.
	return journeys, nil
}
