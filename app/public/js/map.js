var map;
var sidebar;
var allMarkersLayer;
var currentZoomLevel;
var geoJsonLayer;
// Create the map
function createMap() {
  map = L.map("map").setView([51.163361, 10.447683], 7); //Mainz = [49.991756, 8.24414], 15

  startup();
  // Add the OpenStreetMap tiles
  L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
    attribution:
      '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
  }).addTo(map);

  // Add zoom control to the map
  map.removeControl(map.zoomControl);
  L.control.zoom({ position: "bottomright" }).addTo(map);

  allMarkersLayer = L.markerClusterGroup().addTo(map);
  //loadAllAttractions();
}
// Unused function to load all attractions from the go api (unsure how well this works with tens of thousands of attractions)
function loadAllAttractions() {
  fetch(document.location.origin + "/api/attractions", {
    method: "GET",
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    })
    .catch((error) => {
      console.error("There was a problem with the request:", error);
      alert("Failed to load markers. Please try again.");
    });
}

