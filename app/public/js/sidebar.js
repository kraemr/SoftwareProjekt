function toggleSettings() {
    const settingsContainer = document.querySelector(".settings-container");
    settingsContainer.style.display =
        settingsContainer.style.display === "block" ? "none" : "block";
}

// Optional: Close the settings when clicking outside of it
document.addEventListener("click", function (event) {
    const userContainer = document.querySelector(".user-container");
    const settingsContainer = document.querySelector(".settings-container");
    if (
        !userContainer.contains(event.target) &&
        !settingsContainer.contains(event.target)
    ) {
        settingsContainer.style.display = "none";
    }
});
window.addEventListener("resize", function () {
    const sidepanel = document.getElementById("sidepanel-toggle");
    const button = document.getElementById("toggleButton");
    const searchContainer = document.getElementById("search-container");
    // This code will run whenever the window is resized
    if (window.innerWidth <= 768) {
        // Change the sidebar to the small screen layout
        console.log("The window is small");
        sidepanel.style.height = "0";
        sidepanel.style.top = "auto";
        sidepanel.style.bottom = "0";
        sidepanel.style.width = "100%";
        button.style.bottom = "2px";
        button.style.left = "50%";
        searchContainer.style.left = "2px";
    } else {
        // Change the sidebar to the large screen layout
        sidepanel.style.height = "100%";
        sidepanel.style.top = "0";
        sidepanel.style.bottom = "auto";
        sidepanel.style.width = "0";
        button.style.bottom = "50%";
        button.style.left = "2px";
        console.log("The window is large");
        searchContainer.style.left = "2px";
    }
});
// Toggle left sidebar
function toggleSidepanel() {
    const sidepanel = document.getElementById("sidepanel-toggle");
    const button = document.getElementById("toggleButton");
    const searchContainer = document.getElementById("search-container");
    if (window.innerWidth <= 768) {
        // Check if the screen size is small
        if (sidepanel.style.height === "40%") {
            sidepanel.style.height = "0";
            button.style.bottom = "2px";
        } else {
            sidepanel.style.height = "40%";
            button.style.bottom = "41%";
        }
    } else {
        if (sidepanel.style.width === "25%") {
            sidepanel.style.width = "0";
            button.style.left = "2px";
            searchContainer.style.left = "2px";
        } else {
            sidepanel.style.width = "25%";
            button.style.left = "25%";
            searchContainer.style.left = "25%";
        }
    }
}
function openSidepanel() {
    const sidepanel = document.getElementById("sidepanel-toggle");
    const button = document.getElementById("toggleButton");
    const searchContainer = document.getElementById("search-container");
    // Check if the screen size is small
    if (window.innerWidth <= 768) {
        if (sidepanel.style.height = "0") {
            sidepanel.style.height = "40%";
            button.style.bottom = "41%";
        }
    } else {
        if (sidepanel.style.width = "0") {
            sidepanel.style.width = "25%";
            button.style.left = "25%";
            searchContainer.style.left = "25%";
        }
    }
}

function getCategories() {
    var apiUrl = document.location.origin + "/api/categories";
    console.log(apiUrl);
    // Return the promise generated by fetch
    return fetch(apiUrl)
        .then((response) => response.json())
        .catch((error) => {
            console.error("Fehler bei der API-Abfrage:", error);
        });
}
// Function to fill the sidebar categories from the database
function fillCategories() {
    let selectedButton = null;
    const categoriesWrapper = document.getElementById("categories-scroller");
    getCategories().then((categories) => {
        categories.forEach((category) => {
            const button = document.createElement("button");
            button.className = "btn-secondary btn-categories m-2";
            button.innerHTML = category;
            categoriesWrapper.appendChild(button);

            // Add event listener to each button to filter attraction by that category
            button.addEventListener("click", function () {
                // If the clicked button is already selected, unselect it
                if (selectedButton === button) {
                    button.classList.remove("selected");
                    selectedButton = null;
                    // Optionally, you can call a function to handle the unselection case, e.g., showing all markers or clearing the map
                    loadAttractionsByCity(currentCity);
                    currentCategory = null;
                } else {
                    // Remove 'selected' class from the previously selected button
                    if (selectedButton) {
                        selectedButton.classList.remove("selected");
                    }
                    // Add 'selected' class to the clicked button
                    button.classList.add("selected");
                    // Update the selectedButton variable
                    selectedButton = button;
                    loadAttractionsByCategory(category);
                    currentCategory = category;
                }
            });
        });
    });
}

async function showRoute(fromLat, fromLon, toLat, toLon) {  
    var apiUrl = document.location.origin + `/api/public_transport?fromLat=${fromLat}&fromLon=${fromLon}&toLat=${toLat}&toLon=${toLon}`;
  
    fetch(apiUrl)
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        const routeDiv = document.getElementById("routeShowcase");
        routeDiv.innerHTML = '';
  
        data.forEach((journey, index) => {
            const card = document.createElement('div');
            card.className = 'card bg-dark text-light m-2';
  
            const cardHeader = document.createElement('div');
            cardHeader.className = 'card-header';
            cardHeader.textContent = `Option ${index + 1}`;
            cardHeader.onclick = function() {
                const content = this.nextElementSibling;
                content.style.display = content.style.display === 'block' ? 'none' : 'block';
            };
  
            const cardBody = document.createElement('div');
            cardBody.className = 'card-body';
  
            journey.legs.forEach(leg => {
                const depTime = new Date(leg.plannedDeparture).toLocaleString();
                const arrTime = new Date(leg.plannedArrival).toLocaleString();
                const info = document.createElement('p');
                info.innerHTML = `From ${leg.origin.name} to ${leg.destination.name}<br>
                                  Departure: ${depTime}<br>
                                  Arrival: ${arrTime}<br>
                                  Mode: ${leg.line.mode || 'Walking'}${leg.line.name ? ` - ${leg.line.name}` : ''}`;
                cardBody.appendChild(info);
            });

            // Die Ankufstzeit des letzten Legs wird als Ankunftszeit der Route genommen
            const lastLeg = journey.legs[journey.legs.length - 1];
            const arrivalTime = new Date(lastLeg.plannedArrival).toLocaleString();
            const arrivalInfo = document.createElement('p');
            arrivalInfo.innerHTML = `Arrival: ${arrivalTime}`;
            cardHeader.appendChild(arrivalInfo);
  
            card.appendChild(cardHeader);
            card.appendChild(cardBody);
            routeDiv.appendChild(card);
        });
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
    });
}

// Funktion für die Abfrage von "Startort"
function getStartLocation(attLat, attLng) {
    // label "gebe dein Standort ein"
    const routeDiv = document.getElementById("routeShowcase");
    routeDiv.innerHTML = '<p>Enter your start location:</p>';
    
    // Hbox:
    hdiv = document.createElement("div");
    hdiv.className = "hbox";

    inputField = document.createElement("input");
    inputField.type = "text";
    inputField.id = "startLocation";
    inputField.placeholder = "Enter location...";
    hdiv.appendChild(inputField);

    // Button "Use current location"
    button = document.createElement("button");
    button.innerHTML = "Current location";
    button.className = "btn btn-primary";
    button.onclick = function() {
        // Geolocation API
        navigator.geolocation.getCurrentPosition(function(position) {
            const fromLat = position.coords.latitude;
            const fromLon = position.coords.longitude;
    
            // API-Abfrage für die Route
            showRoute(fromLat, fromLon, attLat, attLng);
        });
    };
    hdiv.appendChild(button);

    // Button "Search"
    button = document.createElement("button");
    button.innerHTML = "Search";
    button.className = "btn btn-primary";
    button.onclick = function() {
        // Long/Lat durch nomenatim API abfragen
        const location = inputField.value;
        const apiUrl = `https://nominatim.openstreetmap.org/search?q=${location}&format=json&limit=1`;
        fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            if (data.length > 0) {
                const lat = data[0].lat;
                const lon = data[0].lon;
                // Koordinaten für Start- und Zielort
                const fromLat = lat;
                const fromLon = lon;
                
                // API-Abfrage für die Route
                showRoute(fromLat, fromLon, attLat, attLng);
            }
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });
    };

    routeDiv.appendChild(hdiv);

    routeDiv.appendChild(button);
}

function loadMarkerInfoToSidebar(attractionData) {
    hideSidebarContent();
    openSidepanel();
    const selectedAttractionsInfo = document.getElementById(
        "selectedAttractionsInformation"
    );
    selectedAttractionsInfo.style.display = "block";
    selectedAttractionsInfo.innerHTML = `
<div class="card text-light bg-transparent m-2">
<div class="card-body">
<div id="carouselExampleIndicators" class="carousel slide" data-bs-ride="carousel">
<div class="carousel-inner">
  ${loadCarouselImages()}
</div>
<button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide="prev">
  <span class="carousel-control-prev-icon" aria-hidden="true"></span>
  <span class="visually-hidden">Previous</span>
</button>
<button class="carousel-control-next" type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide="next">
  <span class="carousel-control-next-icon" aria-hidden="true"></span>
  <span class="visually-hidden">Next</span>
</button>
</div>
  <br>
  <h5 class="card-title">${attractionData.title}</h5>
  <p class="card-text">${attractionData.info}</p>
  <p class="card-text">${attractionData.city}</p>
  <p class="card-text">${attractionData.type}</p>
  <p class="card-text">${attractionData.Stars} &#11088;
  <p class="card-text">${attractionData.recommended_count} &#128150;   <span class="favourite-section mt-3" style="position: relative; z-index: 1;">
    <button class="btn btn-warning" id="favouriteButton">
      <i class="fas fa-star"></i> Favourite
    </button>
  </span></p></p>
  
  <button class="btn btn-primary w-100" onclick="getStartLocation(${attractionData.posX}, ${attractionData.posY})">Show Route</button>
  <div id="routeShowcase">
  </div>
  <div class="review-section">
    <h6>Leave a Review:</h6>
    <br>
<div class="star-rating">
    ${[1, 2, 3, 4, 5]
            .map(
                (star) => `
      <span class="star" data-value="${star}">&#9733;</span>
    `
            )
            .join("")}
  </div>
    <textarea class="form-control mt-2" placeholder="Write your review here..."></textarea>
    <button class="btn btn-secondary mt-2">Submit Review</button>
  </div>
</div>
</div>

`;
    // Add event listener for the favourite button
    document.getElementById("favouriteButton").addEventListener("click", function () {
        attractionData.recommended_count += 1;
        fetch(document.location.origin + "/api/attractions", {
            method: "PUT",
            body: JSON.stringify(attractionData),
        })
            .then((response) => response.json())
            .then((data) => {
                console.log(data);
            });
        fetch(document.location.origin + "/api/favorites", {
            method: "POST",
            body: JSON.stringify(attractionData),
        })
            .then((response) => response.json())
            .then((data) => {
                console.log(data);
            });
    });
    // Add event listeners for the star rating
    document.querySelectorAll(".star-rating .star").forEach(star => {
        star.addEventListener("click", function () {
            const rating = this.getAttribute("data-value");
            console.log(`User rated: ${rating} stars`);
            // Update star filling based on rating
            document.querySelectorAll(".star-rating .star").forEach(s => {
                s.classList.remove('filled');
                if (s.getAttribute("data-value") <= rating) {
                    s.classList.add('filled');
                }
            });
            attractionData.stars = rating;
            // Submit the rating
            fetch(document.location.origin + "/api/attractions", {
                method: "PUT",
                body: JSON.stringify(attractionData),
            })
                .then((response) => response.json())
                .then((data) => {
                    console.log(data);
                });
        });
    });
}


function loadCarouselImages() {
    const images = [
        { src: "images/image1.jpg", alt: "First slide" },
        { src: "images/image2.jpg", alt: "Second slide" },
        { src: "images/image3.jpg", alt: "Third slide" },
    ];
    return images
        .map(
            (img, index) => `
      <div class="carousel-item ${index === 0 ? "active" : ""}">
          <img src="${img.src
                }" class="d-block object-fit-cover carousel-images w-100" alt="${img.alt
                }">
      </div>
  `
        )
        .join("");
}
function hideSidebarContent() {
    const selectedAttractionsInfo = document.getElementById(
        "selectedAttractionsInformation"
    );
    const noAttractionSelected = document.getElementById("NoAttractionSelected");
    const showFavourites = document.getElementById("showFavourites");
    selectedAttractionsInfo.style.display = "none";
    noAttractionSelected.style.display = "none";
    showFavourites.style.display = "none";
}
