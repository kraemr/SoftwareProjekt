var map;
var sidebar;
var allMarkersLayer;
var currentZoomLevel;
var breakZoomLevel = 15;
var geoJsonLayer;

function createMap() {
  map = L.map("map").setView([51.163361, 10.447683], 7); //Mainz = [49.991756, 8.24414], 15

  displayGermanyonStartup();
  L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
    attribution:
      '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
  }).addTo(map);

  //Entfernen des Standard-Zoom Bereichs und hinzufügen des neuen Bereichs
  map.removeControl(map.zoomControl);
  L.control.zoom({ position: "bottomright" }).addTo(map);

  allMarkersLayer = L.markerClusterGroup().addTo(map);
  loadAllMarkers();
}
function loadAllMarkers() {
  fetch(document.location.origin + "/api/attractions", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    }
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    })
    .then((data) => {
      placeMarkers(data);
    })
    .catch((error) => {
      console.error("There was a problem with the request:", error);
      alert("Failed to load markers. Please try again.");
    });
}


function placeMarkers(data) {
  for (var elem of data) {
    let marker = createBlueMarker(elem);
    allMarkersLayer.addLayer(marker);
  }
  console.log(data);
}


//Blauer Marker
function createBlueMarker(attraction) {
  var latlng = L.latLng(attraction.posX, attraction.posY);
  var customIcon = L.icon({
    iconUrl: "leaflet/images/marker-icon-2x.png", // Update this path to your actual icon path
    iconSize: [25, 41],
    iconAnchor: [12, 41],
    popupAnchor: [0, -16],
  });

  var marker = L.marker(latlng, {
    icon: customIcon,
    clickable: true,
  });

  // Store attraction data in the marker
  marker.attractionData = attraction;

  // Add click event listener to the marker
  marker.on("click", function () {
    setPopUp(marker.attractionData, marker);
  });

  return marker;
}

function placeMarkers(data) {
  for (var elem of data) {
    let marker = createBlueMarker(elem);
    allMarkersLayer.addLayer(marker);
  }
  console.log(data);
}


function setPopUp(data, marker) {
  var popupContent = `
    <div>
      <strong>City: </strong> ${data.city}<br>
      <strong>Title: </strong> ${data.title}<br>
      <strong>ID: </strong> ${data.Id}<br>
      <strong>Type: </strong> ${data.type}<br>
      <strong>Position X: </strong> ${data.posX}<br>
      <strong>Position Y: </strong> ${data.posY}<br>
      <strong>Info: </strong> ${data.info}<br>
      <strong>Stars: </strong> ${data.Stars}&#11088;<br>
      <strong>Recommended Count: </strong> ${data.recommended_count}
    </div>
  `;
  // Create a new popup instance for the marker
  var popup = L.popup().setContent(popupContent);

  // Bind the popup to the marker
  marker.bindPopup(popup);

  // Open the popup when the marker is clicked
  marker.on("click", function () {
    marker.openPopup();
  });
}
