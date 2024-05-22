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
function createBlueMarker(lat, lng, attractionID) {
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
  marker.attractionID = attractionID;

  // Add click event listener to the marker
  marker.on('click', function () {
    loadPopInformation(marker);
  });

  return marker;
}
function setPopUp(data, marker) {
  var vehicleTrue = "";
  if (data[0].Rad == true) { vehicleTrue += `<strong>Rad: </strong>beteiligt<br>` }
  else if (data[0].Rad == false) { vehicleTrue += `<strong>Rad: </strong>unbeteiligt<br>` }
  if (data[0].PKW == true) { vehicleTrue += `<strong>PKW: </strong>beteiligt<br>` }
  else if (data[0].PKW == false) { vehicleTrue += `<strong>PKW: </strong>unbeteiligt<br>` }
  if (data[0].Fußgänger == true) { vehicleTrue += `<strong>Fußgänger: </strong>beteiligt<br>` }
  else if (data[0].Fußgänger == false) { vehicleTrue += `<strong>Fußgänger: </strong>unbeteiligt<br>` }
  if (data[0].Kraftrad == true) { vehicleTrue += `<strong>Kraftrad: </strong>beteiligt<br>` }
  else if (data[0].Kraftrad == false) { vehicleTrue += `<strong>Kraftrad: </strong>unbeteiligt<br>` }
  if (data[0].Güterkraftfahrzeug == true) { vehicleTrue += `<strong>Güterkraftfahrzeug: </strong>beteiligt<br>` }
  else if (data[0].Güterkraftfahrzeug == false) { vehicleTrue += `<strong>Güterkraftfahrzeug: </strong>unbeteiligt<br>` }
  if (data[0].Sonstige == true) { vehicleTrue += `<strong>Sonstige: </strong>beteiligt<br>` }
  else if (data[0].Sonstige == false) { vehicleTrue += `<strong>Sonstige: </strong>unbeteiligt<br>` }

  var popupContent = `
        <div>
            <strong>Jahr: </strong> ${data[0].Jahr}<br>
            <strong>Unfallkategorie: </strong> ${data[0].Kategorie}<br>
            <strong>Unfallart: </strong> ${data[0].Unfallart}<br>
            <strong>Unfalltyp: </strong> ${data[0].Unfalltyp}<br>
            <strong>Straßenzustand: </strong> ${data[0].Straßenzustand}<br>
            <strong>Lichtverhältnisse: </strong> ${data[0].Lichtverhältnisse}<br>
            ${vehicleTrue}     
        </div>
    `;

  // Set the popup content for the currently clicked marker
  marker.bindPopup(popupContent).openPopup();
}

function loadPopInformation(marker) {
  //Get unfallId from marker
  var unfallId = marker.unfallID;

  //query information from unfallId
  var xhr = new XMLHttpRequest();
  xhr.open('GET', `dbQueries.php?function=loadUnfallById&unfallid=` + unfallId, true);    //Concat get Parameter

  xhr.onload = function () {
    console.log(xhr.responseText); //Handler der auf eine Response wartet, die Anfrage wird erst danach mit xhr.send() aufgerufen
    if (xhr.status === 200) {
      var data = JSON.parse(xhr.responseText);
      //console.log("popup info " + data);
      setPopUp(data, marker);   //call set Filters with DB data
    } else {
      reject('Request failed. Returned status of ' + xhr.status);
    }
  };
  xhr.send();
}

function filterButtonClick() {
  if (currentZoomLevel >= breakZoomLevel) {
    removeAllMarkers();
    loadMarkersInVisibleBounds();
  }
}