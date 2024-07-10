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
// Toggle sidebar size
function toggleSidepanel() {
  const sidepanel = document.getElementById("sidepanel-toggle");
  const button = document.getElementById("toggleButton");
  const searchContainer = document.getElementById("search-container");
  // This is to change the sidebar size between the small and large screen layout
  if (window.innerWidth <= 768) {
    // e.g if the window is small and trhe sidebar is at 40% height, change it to 0
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
          currentCategory = category;
          console.log("Category selected: " + category);
          loadAttractionsByCategory(category);
        }
      });
    });
  });
}
// Function to fill the sidebar with attraction information
function loadMarkerInfoToSidebar(attractionData) {
  hideSidebarContent();
  toggleSidepanel();
  const selectedAttractionsInfo = document.getElementById(
    "selectedAttractionsInformation"
  );
  selectedAttractionsInfo.style.display = "block";
  selectedAttractionsInfo.innerHTML = `
<link rel="stylesheet" href="./css/sidebar.css">
<div class="d-flex justify-content-center mb-2">
  <button class="btn btn-link active-tab" id="overviewLink">Overview</button>
  <button class="btn btn-link" id="reviewsLink">Reviews</button>
</div>
<div class="content-section">
  <div class="card text-light bg-transparent m-2" id="overviewSection">
    <div class="card-body">
      <img src="${attractionData.Img_url}" class="card-img-top" alt="${attractionData.title}">
    </div>
    <br>
    <h5 class="card-title">${attractionData.title}</h5>
    <p class="card-text">${attractionData.info}</p>
    <p class="card-text"><strong>${attractionData.city}</strong>, ${attractionData.Street} ${attractionData.Housenumber}
      <button class="btn btn-primary" id="showRouteBtn" onclick="getStartLocation(${attractionData.posX}, ${attractionData.posY})">Show Route</button>
      <div id="routeShowcase"></div>
    </p>
    <p class="card-text"><strong>Category: </strong>${attractionData.type}</p>
    <p class="card-text">${attractionData.Stars} &#11088;
    <p class="card-text">${attractionData.recommended_count} &#128150;
      <span class="favourite-section mt-3" style="position: relative; z-index: 1;">
        <button class="btn btn-warning" id="favouriteButton">
        <i class="fas fa-star"></i> Favourite
        </button>
      </span>
    </p>
  </div>
  <div class="review-section" id="reviewsSection" style="display: none;">
    <div class="star-rating" id="star-rating-review">
      ${[1, 2, 3, 4, 5]
      .map(
        (star) => `
        <span class="star" data-value="${star}">&#9733;</span>
      `
      )
      .join("")}
    </div>
    <textarea id="reviewAttractionText" class="form-control mt-2" placeholder="Write your review here..."></textarea>
    <button class="btn btn-secondary mt-2" id="submitReview">Submit Review</button>
    <div class="reviews-container mt-3">
      <h5 class="text-white m-2">Reviews</h5>
      <div class="userReviews">`
  fetch(document.location.origin + "/api/users", {
    method: "GET",
  })
    .then((response) => response.json())
    .then((userData) => {
      fetch(document.location.origin + "/api/reviews?attraction_id=" + attractionData.Id, {
        method: "GET",
      })
        .then((response) => response.json())
        .then((reviews) => {
          const reviewsContainer = document.querySelector(".userReviews");
          reviews.forEach((review) => {
            const reviewCard = document.createElement("div");
            reviewCard.classList.add("card", "text-light", "bg-transparent", "m-2");
            var username = review.Username;
            if (review.Username === "") {
              username = "Anonymous";
            }
            console.log(review);
            reviewCard.innerHTML = `
              <div class="card-body">
                <div class="star-rating">
                  ${[1, 2, 3, 4, 5]
                .map(
                  (star) => `
                        <span class="star ${star <= review.Stars ? 'filled' : ''}">&#9733;</span>
                      `
                )
                .join("")}
                </div>
                <p class="card-text">${review.Text}</p>
                <p class="card-text">By ${username} on ${review.Date}
            `;
            if (review.User_id === userData.UserId) {
              const buttonContainer = document.createElement("div");
              const editButton = document.createElement("button");
              editButton.classList.add("btn", "btn-primary", "edit-review", "w-50");
              editButton.innerHTML = "Edit";
              editButton.addEventListener("click", function () {
                // Call the edit review API endpoint here
                console.log("Edit review:", review.Id);
                const updatedReviewText = prompt("Enter the updated review text:");
                console.log(updatedReviewText);
                const reviewData = {
                  Text: updatedReviewText,
                  review_id: review.Id,
                  Username: userData.Username,
                  User_id: userData.UserId,
                  Stars: parseFloat(rating),
                  Date: new Date().toISOString()
                };
                fetch(document.location.origin + "/api/reviews", {
                  method: "PUT",
                  body: JSON.stringify(reviewData),
                  headers: {
                    "Content-Type": "application/json",
                  },
                })
                  .then((response) => response.json())
                  .then((data) => {
                    console.log(data);
                    // Update the review card with the updated text
                    if (data.success) {
                      reviewCard.querySelector(".card-text").textContent = updatedReviewText;
                    }
                  })
                  .catch((error) => {
                    console.error("Error editing review:", error);
                  });
              });
              const deleteButton = document.createElement("button");
              deleteButton.classList.add("btn", "btn-danger", "delete-review", "w-50");
              deleteButton.innerHTML = "Delete";
              deleteButton.addEventListener("click", function () {
                // Call the delete review API endpoint here
                console.log("Delete review:", review.Id)
                fetch(document.location.origin + "/api/reviews?id=" + review.Id, {
                  method: "DELETE",
                })
                  .then((response) => response.json())
                  .then((data) => {
                    console.log(data);
                    // Remove the review card from the DOM
                    if (data.success) {
                      reviewCard.remove();
                    }
                  })
                  .catch((error) => {
                    console.error("Error deleting review:", error);
                  });
              });
              buttonContainer.appendChild(editButton);
              buttonContainer.appendChild(deleteButton);
              reviewCard.appendChild(buttonContainer);
            }
            reviewsContainer.appendChild(reviewCard);
          });
        })
    })
    .catch((error) => {
      console.error("Error fetching reviews:", error);
    });
  `</p>
                </div>
  </div>
</div>

`;
  // Overview and Reviews tabs
  document.getElementById('overviewLink').addEventListener('click', function () {
    document.getElementById('overviewSection').style.display = 'block';
    document.getElementById('reviewsSection').style.display = 'none';
    document.getElementById('overviewLink').classList.add('active-tab');
    document.getElementById('reviewsLink').classList.remove('active-tab');
  });

  document.getElementById('reviewsLink').addEventListener('click', function () {
    document.getElementById('overviewSection').style.display = 'none';
    document.getElementById('reviewsSection').style.display = 'block';
    document.getElementById('reviewsLink').classList.add('active-tab');
    document.getElementById('overviewLink').classList.remove('active-tab');
  });

  // Adding the active-tab class for the initial load
  document.getElementById('overviewLink').classList.add('active-tab');

  // Change the Show Route button to Cancel Button unless its already the cancel button
  const routeButton = document.getElementById("showRouteBtn");
  routeButton.addEventListener("click", function () {
    if (routeButton.innerHTML === "Show Route") {
      routeButton.innerHTML = "Cancel";
    } else {
      routeButton.innerHTML = "Show Route";
      // Clear the route showcase div
      const routeDiv = document.getElementById("routeShowcase");
      routeDiv.innerHTML = '';
    }
  });
  // Check if the current user has already favourited the attraction
  fetch(document.location.origin + "/api/favorites?action=count&attraction_id=" + attractionData.Id, {
    method: "GET",
  })
    .then(response => response.json())
    .then(data => {
      console.log(data);
      const favouriteButton = document.getElementById("favouriteButton");
      if (data.favorite_count > 0) {
        favouriteButton.innerHTML = "Unfavourite";
        favouriteButton.classList.remove("btn-warning");
        favouriteButton.classList.add("btn-danger");
      }
    })
    .catch(error => {
      console.error('Error:', error);
    });


  // Add event listener for the favourite button
  document.getElementById("favouriteButton")
    .addEventListener("click", function () {

      fetch(document.location.origin + "/api/users", {
        method: "GET",
      })
        .then((response) => response.json())
        .then((data) => {
          console.log(data.UserId);
          console.log(attractionData.Id);

          fetch(document.location.origin + "/api/favorites", {
            method: "POST",
            body: JSON.stringify({
              user_id: data.UserId,
              attraction_id: attractionData.Id,
            }),
            headers: {
              "Content-Type": "application/json",
            },
          })
            .then((response) => response.json())
            .then((data) => {
              console.log(data.info);
              if (data.info === "deleted favorite") {
                favouriteButton.innerHTML = "Favourite";
                favouriteButton.classList.add("btn-warning");
                favouriteButton.classList.remove("btn-danger");
                attractionData.recommended_count--;
                loadMarkerInfoToSidebar(attractionData);
              }
              if (data.info === "added favorite") {
                favouriteButton.innerHTML = "Unfavourite";
                favouriteButton.classList.remove("btn-warning");
                favouriteButton.classList.add("btn-danger");
                attractionData.recommended_count++;
                loadMarkerInfoToSidebar(attractionData);
              }
            });
        });
    });
  var rating = 0;
  var review = "";
  // Add event listeners for the star rating
  document.querySelectorAll(".star-rating .star").forEach((star) => {
    star.addEventListener("click", function () {
      rating = this.getAttribute("data-value");
      console.log(`User rated: ${rating} stars`);
      // Update star filling based on rating
      document.querySelectorAll(".star-rating .star").forEach((s) => {
        s.classList.remove("filled");
        if (s.getAttribute("data-value") <= rating) {
          s.classList.add("filled");
        }
      });
      rating = rating;
    });
  });
  // Get the text review from the textarea
  const SubmitReviewButton = document.getElementById("submitReview");
  SubmitReviewButton.addEventListener("click", function () {
    const reviewText = document.getElementById("reviewAttractionText").value;
    console.log(reviewText, rating);
    // if the user has rated the attraction and input text, send the review
    if (rating > 0) {
      fetch(document.location.origin + "/api/users", {
        method: "GET",
      })
        .then((response) => response.json())
        .then((userData) => {
          console.log(userData.Username);
          const reviewData = {
            Text: reviewText,
            Attraction_id: attractionData.Id,
            Username: userData.Username,
            User_id: userData.UserId,
            Stars: parseFloat(rating),
            Date: new Date().toISOString()
          };
          fetch(document.location.origin + "/api/reviews", {
            method: "POST",
            body: JSON.stringify(reviewData),
            headers: {
              "Content-Type": "application/json",
            },
          })
            .then((response) => response.json())
            .then((data) => {
              if (data.success) {
                console.log('Review added:', data);
                // Clear the review textarea
                document.getElementById("reviewAttractionText").value = "";
                // Update the reviews section
                loadMarkerInfoToSidebar(attractionData);
                // Open the reviews tab
                document.getElementById('reviewsLink').click();
              }
            });
        }
        );
    }
  }
  );
}

// Get Route from locationA to locationB
async function showRoute(fromLat, fromLon, toLat, toLon) {
  // Create a new route planning div
  const routePlanningDiv = document.createElement("div");
  routePlanningDiv.id = "routePlanningDiv";
  routePlanningDiv.innerHTML = "Route Planning Div";
  document.body.appendChild(routePlanningDiv);
  var apiUrl =
    document.location.origin +
    `/api/public_transport?fromLat=${fromLat}&fromLon=${fromLon}&toLat=${toLat}&toLon=${toLon}`;

  // Call REST API to get the route
  fetch(apiUrl)
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    })
    .then(data => {
      const routeDiv = document.getElementById("routeShowcase");
      // Clear the route showcase div before adding new content
      routeDiv.innerHTML = '';

      // Add "option" for each journey
      data.forEach((journey, index) => {
        const card = document.createElement('div');
        card.className = 'card bg-dark text-light m-2';

        const cardHeader = document.createElement('div');
        cardHeader.className = 'card-header';
        cardHeader.textContent = `Option ${index + 1}`;

        // Add event listener to the card header to show/hide the content for better visibility
        cardHeader.onclick = function () {
          const content = this.nextElementSibling;
          content.style.display = content.style.display === 'block' ? 'none' : 'block';
        };

        const cardBody = document.createElement('div');
        cardBody.className = 'card-body';

        // Add information for each leg ("step") of the journey
        journey.legs.forEach(leg => {
          const depTime = new Date(leg.plannedDeparture).toLocaleString();
          const arrTime = new Date(leg.plannedArrival).toLocaleString();
          const info = document.createElement('p');
          info.innerHTML = `From ${leg.origin.name} to ${leg.destination.name}<br>
                                  Departure: ${depTime}<br>
                                  Arrival: ${arrTime}<br>
                                  Mode: ${leg.line.mode || "Walking"}${leg.line.name ? ` - ${leg.line.name}` : ""
            }`;
          cardBody.appendChild(info);

          // Add "Show on Map" button for each leg
          const showOnMapBtn = document.createElement('button');
          showOnMapBtn.className = 'btn btn-secondary m-2';
          showOnMapBtn.innerHTML = 'Zoom on Map';
          showOnMapBtn.onclick = function () {
            publicTransportMarker(leg);
          };

          cardBody.appendChild(showOnMapBtn);

          // Add horizontal line after each leg
          const hr = document.createElement('hr');
          cardBody.appendChild(hr);
        });

        // The arrival time of the last leg is taken as the arrival time of the option
        const lastLeg = journey.legs[journey.legs.length - 1];
        const arrivalTime = new Date(lastLeg.plannedArrival).toLocaleString();
        const arrivalInfo = document.createElement("p");
        arrivalInfo.innerHTML = `Arrival: ${arrivalTime}`;
        cardHeader.appendChild(arrivalInfo);

        // Add the card to the route showcase div
        card.appendChild(cardHeader);
        card.appendChild(cardBody);
        routeDiv.appendChild(card);
      });
    })
    .catch((error) => {
      console.error("There was a problem with the fetch operation:", error);
    });
}

// Get the attraction location as parameters on call
function getStartLocation(attLat, attLng) {
  const routeDiv = document.getElementById("routeShowcase");
  hdiv = document.createElement("div");
  hdiv.className = "hbox";

  inputField = document.createElement("input");
  inputField.type = "text";
  inputField.id = "startLocation";
  inputField.placeholder = "Enter Start location...";
  hdiv.appendChild(inputField);

  // Button "Use current location"
  button = document.createElement("button");
  button.innerHTML = "Current location";
  button.className = "btn btn-primary m-1";
  button.onclick = function () {
    // Geolocation API
    navigator.geolocation.getCurrentPosition(function (position) {
      const fromLat = position.coords.latitude;
      const fromLon = position.coords.longitude;

      // API request for the route
      showRoute(fromLat, fromLon, attLat, attLng);
    });
  };
  hdiv.appendChild(button);

  // Add search button
  button = document.createElement("button");
  button.innerHTML = "Search";
  button.className = "btn btn-primary m-1";
  button.onclick = function () {
    // Retrieve long/lat of provided location through nomenatim API
    const location = inputField.value;
    const apiUrl = `https://nominatim.openstreetmap.org/search?q=${location}&format=json&limit=1`;
    fetch(apiUrl)
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        if (data.length > 0) {
          const lat = data[0].lat;
          const lon = data[0].lon;
          // Coordinates for start and destination
          const fromLat = lat;
          const fromLon = lon;

          // API request for the route
          showRoute(fromLat, fromLon, attLat, attLng);
        }
      })
      .catch((error) => {
        console.error("There was a problem with the fetch operation:", error);
      });
  };
  // Add elements to the route showcase div
  routeDiv.appendChild(hdiv);
  routeDiv.appendChild(button);
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
