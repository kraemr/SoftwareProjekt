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
    let marker = createBlueMarker(elem.posX, elem.posY, elem.Id);
    allMarkersLayer.addLayer(marker);
  }
  console.log(data);
}

//Blauer Marker
function createBlueMarker(lat, lng, attractionID) {
  var latlng = L.latLng(lat, lng);
  var customIcon = L.icon({
    iconUrl: "leaflet/images/marker-icon-2x.png",
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
  fetch(document.location.origin + `/api/attractions/${marker.attractionID}`)
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    })
    .then((data) => {
      setPopUp(data, marker);
    })
    .catch((error) => {
      console.error("There was a problem with the request:", error);
    });
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
      <strong>Stars: </strong> ${data.Stars}<br>
      <strong>Recommended Count: </strong> ${data.recommended_count}
    </div>
  `;

  marker.bindPopup(popupContent).openPopup();
}
