package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"src/public_transport"
	"strings"
	"testing"
)

// TestFetchFullRouteLongLat testet die FetchFullRouteLongLat Funktion.
func TestFetchFullRouteLongLat(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "locations/nearby") {
			resp := []public_transport.Location{
				{ID: "location-id"},
			}
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
		} else if strings.Contains(r.URL.Path, "journeys") {
			resp := struct {
				Journeys []public_transport.Journey `json:"journeys"`
			}{
				Journeys: []public_transport.Journey{
					{Type: "journey"},
				},
			}
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
		}
	}))
	defer testServer.Close()

	journeys, err := public_transport.FetchFullRouteLongLat(49.9179102, 8.3430285, 49.987809, 8.2272517)
	if err != nil {
		t.Fatalf("FetchFullRouteLongLat failed: %v", err)
	}
	if len(journeys) == 0 {
		t.Errorf("Expected to find journeys, found none")
	}
}
