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