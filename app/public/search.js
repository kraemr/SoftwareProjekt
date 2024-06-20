var currentCity;
function searchBox() {
    // Read Input Box
    var query = document.getElementById('search-input').value;
    searchLocation(query, true);
}
function searchLocation(query, clearSearchInput) {
    // Check for empty string in query
    if (query.trim() === "") {
        alert("Bitte geben Sie einen Suchbegriff ein.");
        return;
    }

    // Regex for GPS coordinates
    var coordinatesRegex = /^([-+]?\d{1,2}(?:[.,]\d+)?)[\s,]+([-+]?\d{1,3}(?:[.,]\d+)?)$/;
    var match = query.match(coordinatesRegex);

    if (match) {
        var lat = parseFloat(match[1].replace(',', '.')); // Breitengrad
        var lon = parseFloat(match[2].replace(',', '.')); // Längengrad
        map.flyTo([lat, lon], 15)
        return;
    }

    // API-URL
    var apiUrl = 'https://nominatim.openstreetmap.org/search.php?q=' + encodeURIComponent(query) + '&polygon_geojson=1&format=geojson&limit=1&countrycodes=de';
    console.log(apiUrl);

    // Ausführen der API-Abfrage
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            if (query != "Deutschland") {
                loadAttractionsByCity(query);
                currentCity = query;
            }
            updateGeoJsonLayer(data);
            if (clearSearchInput) {
                document.getElementById('search-input').value = "";
            }
        })
        .catch(error => {
            console.error('Fehler bei der API-Abfrage:', error);
        });
}
// Place markers by category and city
function loadAttractionsByCategory(category) {
    var apiUrl = document.location.origin + '/api/attractions?category=' + encodeURIComponent(category);
    console.log(apiUrl);

    // Ausführen der API-Abfrage
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            console.log(currentCity);
            // Löschen aller Marker
            allMarkersLayer.clearLayers();
            data = data.filter(attraction => attraction.city === currentCity);
            // Hinzufügen der neuen Marker
            placeMarkers(data);
        })
        .catch(error => {
            console.error('Fehler bei der API-Abfrage:', error);
        });
}

function loadAttractionsByCity(city) {
    var apiUrl = document.location.origin + '/api/attractions?city=' + encodeURIComponent(city);
    console.log(apiUrl);

    // Ausführen der API-Abfrage
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            console.log(data);
            // Löschen aller Marker
            allMarkersLayer.clearLayers();
            // Hinzufügen der neuen Marker
            placeMarkers(data);
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
    // Check if geolocation is supported by the browser
    if ("geolocation" in navigator) {
        // Prompt user for permission to access their location
        navigator.geolocation.getCurrentPosition(
            // Success callback function
            (position) => {
                // Get the user's latitude and longitude coordinates
                const lat = position.coords.latitude;
                const lng = position.coords.longitude;

                // Display the user location on the map
                console.log(`Latitude: ${lat}, longitude: ${lng}`);
                var getUserLocationByApi = 'https://nominatim.openstreetmap.org/reverse?format=geojson&polygon_geojson=1&format=geojson&limit=1&lat=' + lat + '&lon=' + lng;
                fetch(getUserLocationByApi)
                    .then(response => response.json())
                    .then(data => {
                        // get the city name
                        console.log(data);
                        var city
                        if (data.features[0].properties.address.city == undefined) {
                            city = data.features[0].properties.address.municipality;
                        } else {
                            city = data.features[0].properties.address.city;
                        }
                        console.log(city);
                        searchLocation(city, false);
                    })
                // Create a custom icon for the user's location marker
                var userLocationIcon = L.icon({
                    iconUrl: 'leaflet/images/marker-icon-red.png',
                    iconSize: [25, 41], // Customize the size as needed
                    iconAnchor: [12, 41],
                    popupAnchor: [0, -16],
                });

                // Create the marker with the custom icon
                var userLocationMarker = L.marker([lat, lng], {
                    icon: userLocationIcon,
                    clickable: true,
                }).addTo(map);

                // Set the popup content for the user's location marker
                var popupContent = `
            <div>
                <strong>Your Location</strong><br>
                Latitude: ${lat}<br>
                Longitude: ${lng}
            </div>
        `;
                userLocationMarker.bindPopup(popupContent).openPopup();
            },
            // Error callback function
            (error) => {
                // Handle errors, e.g. user denied location sharing permissions
                console.error("Error getting user location:", error);
                searchLocation("Deutschland", false);
            }

        );
    } else {
        // Geolocation is not supported by the browser
        console.error("Geolocation is not supported by this browser.");
        searchLocation("Deutschland", false);
    }
}

