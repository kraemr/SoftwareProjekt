var map;
var sidebar;
var allMarkersLayer;
var currentZoomLevel;
var breakZoomLevel = 15;
var geoJsonLayer;

function createMap() {
  map = L.map("map").setView([51.163361, 10.447683], 7); //Mainz = [49.991756, 8.24414], 15

  displayGermanyonStartup();
  loadAllMarkers();
  L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
    attribution:
      '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
  }).addTo(map);

  //Entfernen des Standard-Zoom Bereichs und hinzufügen des neuen Bereichs
  map.removeControl(map.zoomControl);
  L.control.zoom({ position: "bottomright" }).addTo(map);

  allMarkersLayer = L.layerGroup().addTo(map);
}
function loadAllMarkers() {
  fetch("/api/attractions")
    .then((response) => response.json())
    .then((data) => {
      placeMarkers(data);
    })
    .catch((error) => {
      console.error("Error fetching attractions:", error);
    });
}

function placeMarkers(data) {
  for (var elem of data) {
    let marker = createBlueMarker(elem.PosX, elem.PosY, elem.id);
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
  //Get unfallId from marker
  var attractionID = marker.attractionID;

  //query information from unfallId
  var xhr = new XMLHttpRequest();
  // GET /api/attractions RETURNS JSON
  xhr.open("GET", "/api/attractions/" + attractionID, true);

  xhr.onload = function () {
    console.log(xhr.responseText); //Handler der auf eine Response wartet, die Anfrage wird erst danach mit xhr.send() aufgerufen
    if (xhr.status === 200) {
      var data = JSON.parse(xhr.responseText);
      console.log("popup info " + data);
      setPopUp(data, marker); //call set Filters with DB data
    } else {
      reject("Request failed. Returned status of " + xhr.status);
    }
  };
  xhr.send();
}
