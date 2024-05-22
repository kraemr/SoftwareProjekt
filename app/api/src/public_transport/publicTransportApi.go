package public_transport;

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "time"
)

const baseURL = "https://v6.db.transport.rest"

type Location struct {
    Type     string `json:"type"`
    ID       string `json:"id"`
    Name     string `json:"name"`
    Location struct {
        Latitude  float64 `json:"latitude"`
        Longitude float64 `json:"longitude"`
    } `json:"location"`
}

type Route struct {
    Type string `json:"type"`
    Legs []struct {
        TripID        string `json:"tripId"`
        Direction     string `json:"direction"`
        Line          struct {
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

func fetchLocation(query string) (Location, error) {
    var location Location
    resp, err := http.Get(fmt.Sprintf("%s/locations?query=%s&results=1", baseURL, url.QueryEscape(query)))
    if err != nil {
        return location, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return location, err
    }
    var locations []Location
    err = json.Unmarshal(body, &locations)
    if err != nil {
        return location, err
    }
    if len(locations) > 0 {
        location = locations[0]
    }
    return location, nil
}

func fetchRoutes(fromID, toID string) ([]Route, error) {
    var routes []Route
    departureTime := time.Now().Add(24 * time.Hour).Format(time.RFC3339)
    resp, err := http.Get(fmt.Sprintf("%s/journeys?from=%s&to=%s&departure=%s&results=1", baseURL, fromID, toID, url.QueryEscape(departureTime)))
    if err != nil {
        return routes, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return routes, err
    }
    var data struct {
        Routes []Route `json:"journeys"`
    }
    err = json.Unmarshal(body, &data)
    if err != nil {
        return routes, err
    }
    routes = data.Routes
    return routes, nil
}

func main() {
    fromLocation, err := fetchLocation("Mainz HBF")
    if err != nil {
        fmt.Println("Error fetching from location:", err)
        return
    }
    toLocation, err := fetchLocation("Lucy-Hillebrand-Straße, Mainz")
    if err != nil {
        fmt.Println("Error fetching to location:", err)
        return
    }
    routes, err := fetchRoutess(fromLocation.ID, toLocation.ID)
    if err != nil {
        fmt.Println("Error fetching routes:", err)
        return
    }
    for _, route := range routes {
        fmt.Printf("Route from %s to %s:\n", fromLocation.Name, toLocation.Name)
        for _, leg := range route.Legs {
            fmt.Printf("  - From %s to %s\n", leg.Origin.Name, leg.Destination.Name)
            fmt.Printf("    Departure: %s (planned: %s)\n", leg.Departure, leg.PlannedDeparture)
            fmt.Printf("    Arrival: %s (planned: %s)\n", leg.Arrival, leg.PlannedArrival)
        }
    }
}
