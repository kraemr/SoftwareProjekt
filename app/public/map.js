var map;
var sidebar;
var allMarkersLayer;
var currentZoomLevel;
var breakZoomLevel = 15;
var geoJsonLayer;

function createMap() {
  map = L.map('map').setView([51.163361, 10.447683], 7); //Mainz = [49.991756, 8.24414], 15

  displayGermanyonStartup();

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
  }).addTo(map);

  //Entfernen des Standard-Zoom Bereichs und hinzufügen des neuen Bereichs
  map.removeControl(map.zoomControl);
  L.control.zoom({ position: 'bottomright' }).addTo(map);

  allMarkersLayer = L.layerGroup().addTo(map);
}
function placeMarkers(data) {
  for (var elem of data) {
    let marker = createBlueMarker(elem.YGCSWGS84, elem.XGCSWGS84, elem.id);
    allMarkersLayer.addLayer(marker);
  }
  console.log(data);
}


//Blauer Marker
function createBlueMarker(lat, lng, unfallId) {
  var latlng = L.latLng(lat, lng);
  var customIcon = L.icon({
    iconUrl: ".leaflet/images/marker-icon-2x-blue.png",
    iconSize: [25, 41], // Größe des Icons
    iconAnchor: [12, 41], // Position des Ankers relativ zur Mitte des Icons
    popupAnchor: [0, -16] // Position des Popups relativ zur Mitte des Icons
  });

  var marker = L.marker(latlng, {
    icon: customIcon,
    clickable: true
  });
  marker.unfallID = unfallId;

  // Add click event listener to the marker
  marker.on('click', function () {
    loadPopInformation(marker);
  });

  return marker;
}