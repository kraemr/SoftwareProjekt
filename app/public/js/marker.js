function placeMarkers(data) {
  for (var elem of data) {
    // Fetch the stars for each attraction loaded
    fetch(`/api/reviews?action=stars&attraction_id=${elem.Id}`)
      .then((response) => response.json())
      // If the array is less than 1 long, soft fail
      .catch((error) => {
        console.error("Error:", error);
        return [];
      })
      .then((data) => {
        elem.Stars = data.stars;
      });
      let marker = createBlueMarker(elem);
      allMarkersLayer.addLayer(marker);
  }
}

// Stop marker
function publicTransportMarker(leg) {
  // Remove all route markers from the map
  const markers = allMarkersLayer.getLayers();
  markers.forEach((marker) => {
    if (marker.options.isRouteMarker) {
      allMarkersLayer.removeLayer(marker);
    }
  });

  // Create a new marker
  var latlng = L.latLng(leg.destination.location.latitude, leg.destination.location.longitude);
  let marker = L.marker(latlng, { isRouteMarker: true });

  // Custom icon for the marker
  var customIcon = L.icon({
    iconUrl: "leaflet/images/marker-icon-2x-red.png", // Update this path to your actual icon path
    iconSize: [25, 41],
    iconAnchor: [12, 41],
    popupAnchor: [0, -16],
  });
  marker.setIcon(customIcon);

  // Convert planned arrival time to a more readable format
  let plannedArrivalTime = new Date(leg.plannedArrival);

  // Add popup with destination name and planned arrival time to the marker
  marker.bindPopup(`${leg.destination.name}<br>
    <strong>Planned arrival Time:</strong> ${plannedArrivalTime.toLocaleTimeString()}`);
  allMarkersLayer.addLayer(marker);

  // Zoom to the marker
  map.flyTo(marker.getLatLng(), 17);
}

//Blue Marker
function createBlueMarker(attraction) {
  var latlng = L.latLng(attraction.posX, attraction.posY);
  var customIcon = L.icon({
    iconUrl: "leaflet/images/marker-icon-2x.png", // Update this path to your actual icon path
    iconSize: [25, 41],
    iconAnchor: [12, 41],
    popupAnchor: [0, -16],
  });

  var marker = L.marker(latlng, {
    clickable: true,
  });

  // Store attraction data in the marker
  marker.attractionData = attraction;

  // Add click event listener to the marker
  marker.on("click", function () {
    setPopUp(marker.attractionData, marker);
    marker.openPopup();
    loadMarkerInfoToSidebar(marker.attractionData);
  });

  return marker;
}
// Function to set the popup content for each marker
function setPopUp(data, marker) {
  var popupContent = `
      <div>
        <strong>Title: </strong> ${data.title}<br>
        <strong>City: </strong> ${data.city}<br>
        <strong>Category: </strong> ${data.type}<br>
        <strong>Description: </strong> ${data.info}<br>
        <strong>Recommended Count: </strong> ${data.recommended_count}<br>
        <strong>Stars: </strong> ${data.Stars}&#11088;
        ${data.image ? `<img src="${data.image}" alt="Image" style="width: 100%; height: auto;">` : ""}
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
function showAttractionOnMap(attraction) {
  // remove all markers from the map
  allMarkersLayer.clearLayers();
  var marker = createBlueMarker(attraction);
  allMarkersLayer.addLayer(marker);
  // zoom to marker
  map.flyTo(marker.getLatLng(), 15);
}