function searchLocation() {
    // Lesen des Inhalts der Input-Box
    var query = document.getElementById('search-input').value;

    // Überprüfen, ob die Suchanfrage nicht leer ist
    if (query.trim() === "") {
        alert("Bitte geben Sie einen Suchbegriff ein.");
        return;
    }

    // Regex-Ausdruck für GPS-Koordinaten-Überprüfung
    var coordinatesRegex = /^([-+]?\d{1,2}(?:[.,]\d+)?)[\s,]+([-+]?\d{1,3}(?:[.,]\d+)?)$/;
    var match = query.match(coordinatesRegex);

    if (match) {
        var lat = parseFloat(match[1].replace(',', '.')); // Breitengrad
        var lon = parseFloat(match[2].replace(',', '.')); // Längengrad
        map.flyTo([lat, lon], 15)
        return;
    }

    // Erstellen der API-URL
    var apiUrl = 'https://nominatim.openstreetmap.org/search.php?q=' + encodeURIComponent(query) + '&polygon_geojson=1&format=geojson&limit=1&countrycodes=de';
    console.log(apiUrl);

    // Ausführen der API-Abfrage
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            updateGeoJsonLayer(data);
            document.getElementById('search-input').value = "";
        })
        .catch(error => {
            console.error('Fehler bei der API-Abfrage:', error);
        });
}

// Funktion zum Aktualisieren der GeoJSON-Schicht und Anpassen der Karte
function updateGeoJsonLayer(data) {
    // GeoJSON-Schicht erstellen oder aktualisieren
    if (geoJsonLayer) {
        geoJsonLayer.clearLayers(); // Falls die Schicht bereits vorhanden ist, lösche sie
    }
    geoJsonLayer = L.geoJSON(data, {
        style: {
            fillColor: 'lightblue',
            fillOpacity: 0.4,
            color: 'blue',
            weight: 1
        }
    }).addTo(map);

    // Anpassen der Karte, um die GeoJSON-Schicht einzuschließen
    map.flyToBounds(geoJsonLayer.getBounds());
}

function displayGermanyonStartup() {

    var apiUrl = 'https://nominatim.openstreetmap.org/search.php?q=Deutschland&polygon_geojson=1&format=geojson&limit=1&countrycodes=de';
    console.log(apiUrl);
    // Check if geolocation is supported by the browser
    if ("geolocation" in navigator) {
        // Prompt user for permission to access their location
        navigator.geolocation.getCurrentPosition(
            // Success callback function
            (position) => {
                // Get the user's latitude and longitude coordinates
                const lat = position.coords.latitude;
                const lng = position.coords.longitude;

                // Do something with the location data, e.g. display on a map
                console.log(`Latitude: ${lat}, longitude: ${lng}`);
                newApiUrl = 'https://nominatim.openstreetmap.org/reverse?format=geojson&lat=' + lat + '&lon=' + lng;
                executeAPICall(newApiUrl);
            },
            // Error callback function
            (error) => {
                // Handle errors, e.g. user denied location sharing permissions
                console.error("Error getting user location:", error);
                executeAPICall(apiUrl);
            }
        );
    } else {
        // Geolocation is not supported by the browser
        console.error("Geolocation is not supported by this browser.");
        executeAPICall(apiUrl);
    }
    // Ausführen der API-Abfrage
}
function executeAPICall(apiUrl) {
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            // GeoJSON-Schicht erstellen oder aktualisieren
            if (geoJsonLayer) {
                geoJsonLayer.clearLayers(); // Falls die Schicht bereits vorhanden ist, lösche sie
            }
            geoJsonLayer = L.geoJSON(data, {
                style: {
                    fillColor: 'lightblue',
                    fillOpacity: 0,
                    color: 'black',
                    weight: 2
                }
            }).addTo(map);

            // Zoom to the nearest city
            const cityCoordinates = data.features[0].geometry.coordinates;
            const cityLatLng = L.latLng(cityCoordinates[1], cityCoordinates[0]);
            map.setView(cityLatLng, 10); // Adjust the zoom level as needed

        })
        .catch(error => {
            console.error('Fehler bei der API-Abfrage:', error);
        });
}

