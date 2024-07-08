function placeMarkers(data) {
    for (var elem of data) {
      let marker = createBlueMarker(elem);
      allMarkersLayer.addLayer(marker);
    }
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
    map.fitBounds(allMarkersLayer.getBounds());
    allMarkersLayer.clearLayers();
    var marker = createBlueMarker(attraction);
    allMarkersLayer.addLayer(marker);
    // zoom to marker
    map.flyTo(marker.getLatLng(), 15);
  }