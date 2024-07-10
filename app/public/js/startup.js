// Called from index.html on startup, the name isnt accurate anymore
function startup() {
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