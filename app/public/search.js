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

//Input-Trigger
function createSearchTrigger() {
    var input = document.getElementById("search-input");

    input.addEventListener("keypress", function (event) {
        if (event.key === "Enter") {
            event.preventDefault();
            document.getElementById("search-button").click();
        }
    });
}

function displayGermanyonStartup() {

    var apiUrl = 'https://nominatim.openstreetmap.org/search.php?q=Deutschland&polygon_geojson=1&format=geojson&limit=1&countrycodes=de';
    console.log(apiUrl);

    // Ausführen der API-Abfrage
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

            // Anpassen der Karte, um die GeoJSON-Schicht einzuschließen
            map.flyToBounds(geoJsonLayer.getBounds());
        })
        .catch(error => {
            console.error('Fehler bei der API-Abfrage:', error);
        });
}