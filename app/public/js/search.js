var currentCity;
var currentCategory;
function searchBox() {
  // Read Input Box
  var query = document.getElementById("search-input").value;
  searchLocation(query, true);
}
function searchLocation(query, clearSearchInput) {
  // Check for empty string in query
  if (query.trim() === "") {
    alert("Bitte geben Sie einen Suchbegriff ein.");
    return;
  }

  // Regex for GPS coordinates
  var coordinatesRegex =
    /^([-+]?\d{1,2}(?:[.,]\d+)?)[\s,]+([-+]?\d{1,3}(?:[.,]\d+)?)$/;
  var match = query.match(coordinatesRegex);

  if (match) {
    var lat = parseFloat(match[1].replace(",", ".")); // Breitengrad
    var lon = parseFloat(match[2].replace(",", ".")); // Längengrad
    map.flyTo([lat, lon], 15);
    return;
  }

  // API-URL
  var apiUrl =
    "https://nominatim.openstreetmap.org/search.php?q=" +
    encodeURIComponent(query) +
    "&polygon_geojson=1&format=geojson&limit=1&countrycodes=de";
  console.log(apiUrl);

  // Ausführen der API-Abfrage
  fetch(apiUrl)
    .then((response) => response.json())
    .then((data) => {
      if (query != "Deutschland") {
        loadAttractionsByCity(query);
        currentCity = query;
      }
      updateGeoJsonLayer(data);
      if (clearSearchInput) {
        document.getElementById("search-input").value = "";
      }
    })
    .catch((error) => {
      console.error("Fehler bei der API-Abfrage:", error);
    });
}
// Place markers by category and city
function loadAttractionsByCategory(category) {
  var apiUrl =
    document.location.origin +
    "/api/attractions?category=" +
    encodeURIComponent(category) +
    "&city=" +
    encodeURIComponent(city);
  console.log(apiUrl);

  // Ausführen der API-Abfrage
  fetch(apiUrl)
    .then((response) => response.json())
    .then((data) => {
      console.log(currentCity);
      // Löschen aller Marker
      allMarkersLayer.clearLayers();
      data = data.filter((attraction) => attraction.city === currentCity);
      // Hinzufügen der neuen Marker
      placeMarkers(data);
    })
    .catch((error) => {
      console.error("Fehler bei der API-Abfrage:", error);
    });
}

function loadAttractionsByCity(city) {
  var apiUrl =
    document.location.origin +
    "/api/attractions?city=" +
    encodeURIComponent(city);
  console.log(apiUrl);

   // Ausführen der API-Abfrage
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            console.log(data);
            // Löschen aller Marker
            allMarkersLayer.clearLayers();
            // Hinzufgen der neuen Marker
            if (currentCategory) {
                data = data.filter(attraction => attraction.type === currentCategory);
            }
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
      fillColor: "lightblue",
      fillOpacity: 0.4,
      color: "blue",
      weight: 1,
    },
  }).addTo(map);

  // Anpassen der Karte, um die GeoJSON-Schicht einzuschließen
  map.flyToBounds(geoJsonLayer.getBounds());
}
