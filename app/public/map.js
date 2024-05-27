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

  allMarkersLayer = L.layerGroup().addTo(map);
  loadAllMarkers();
}
function loadAllMarkers() {
  // Construct the JSON object
  var data;
  jsonData = JSON.stringify(data);
    fetch(document.location.origin + "/api/attractions", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      body: jsonData,
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    })
    .then((data) => {
      console.log(data);
    })
    .catch((error) => {
      // Handle errors
      console.error(
        "There was a problem with the login request:",
        error
      );
      alert("Login failed. Please try again.");
    });
}

function placeMarkers(data) {
  for (var elem of data) {
    let marker = createBlueMarker(elem.posx, elem.posy, elem.id);
    allMarkersLayer.addLayer(marker);
  }
  console.log(data);
}

//Blauer Marker
function createBlueMarker(lat, lng, attractionID) {
  var latlng = L.latLng(lat, lng);
  var customIcon = L.icon({
    iconUrl: ".leaflet/images/marker-icon-2x-blue.png",
    iconSize: [25, 41], // Größe des Icons
    iconAnchor: [12, 41], // Position des Ankers relativ zur Mitte des Icons
    popupAnchor: [0, -16], // Position des Popups relativ zur Mitte des Icons
  });

  var marker = L.marker(latlng, {
    icon: customIcon,
    clickable: true,
  });
  marker.attractionID = attractionID;

  // Add click event listener to the marker
  marker.on("click", function () {
    loadPopInformation(marker);
  });

  return marker;
}
function setPopUp(data, marker) {
  var popupContent = `
        <div>
            <strong>City: </strong> ${data[0].city}<br>
            <strong>Title: </strong> ${data[0].title}<br>
            <strong>ID: </strong> ${data[0].id}<br>
            <strong>Category: </strong> ${data[0].category}<br>
            <strong>Position X: </strong> ${data[0].posx}<br>
            <strong>Position Y: </strong> ${data[0].posy}<br>
            ${vehicleTrue}     
        </div>
    `;

  // Set the popup content for the currently clicked marker
  marker.bindPopup(popupContent).openPopup();
}

function loadPopInformation(marker) {
}
