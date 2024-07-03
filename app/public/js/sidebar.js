function toggleSettings() {
    const settingsContainer = document.querySelector('.settings-container');
    settingsContainer.style.display = (settingsContainer.style.display === 'block') ? 'none' : 'block';
}

// Optional: Close the settings when clicking outside of it
document.addEventListener('click', function (event) {
    const userContainer = document.querySelector('.user-container');
    const settingsContainer = document.querySelector('.settings-container');
    if (!userContainer.contains(event.target) && !settingsContainer.contains(event.target)) {
        settingsContainer.style.display = 'none';
    }
});
window.addEventListener('resize', function () {
    const sidepanel = document.getElementById("sidepanel-toggle");
    const button = document.getElementById("toggleButton");
    // This code will run whenever the window is resized
    if (window.innerWidth <= 768) {
        // Change the sidebar to the small screen layout
        console.log('The window is small');
        sidepanel.style.height = "0";
        sidepanel.style.top = "auto";
        sidepanel.style.bottom = "0";
        sidepanel.style.width = "100%";
        button.style.bottom = "2px";
        button.style.left = "50%";
    } else {
        // Change the sidebar to the large screen layout
        sidepanel.style.height = "100%";
        sidepanel.style.top = "0";
        sidepanel.style.bottom = "auto";
        sidepanel.style.width = "0";
        button.style.bottom = "50%";
        button.style.left = "2px";
        console.log('The window is large');
    }
});
// Toggle left sidebar
function toggleSidepanel() {
    const sidepanel = document.getElementById("sidepanel-toggle");
    const button = document.getElementById("toggleButton");
    if (window.innerWidth <= 768) { // Check if the screen size is small
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
        } else {
            sidepanel.style.width = "25%";
            button.style.left = "25%";
        }
    }
}
function getCategories() {
    var apiUrl = document.location.origin + '/api/categories';
    console.log(apiUrl);

    // Ausführen der API-Abfrage
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            fillCategories(data);
        })
        .catch(error => {
            console.error('Fehler bei der API-Abfrage:', error);
        });
    // Function to fill the sidebar categories from the database
    function fillCategories(data) {
        const categoriesWrapper = document.getElementById("categories-scroller");

        // Variable to keep track of the currently selected button
        let selectedButton = null;

        data.forEach(category => {
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

    }
}
function loadMarkerInfoToSidebar(attractionData) {
    const selectedAttractionsInfo = document.getElementById("selectedAttractionsInformation");
    const noAttractionSelected = document.getElementById("NoAttractionSelected");
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
    <h5 class="card-title">${attractionData.title}</h5>
    <p class="card-text">${attractionData.info}</p>
    <p class="card-text">${attractionData.city}</p>
    <p class="card-text">${attractionData.type}</p>
    <p class="card-text">${attractionData.Stars}</p>
    <p class="card-text">${attractionData.recommended_count}</p>
    
    <a href="#" class="btn btn-primary">Route Planen</a>
    <div class="review-section">
      <h6>Leave a Review:</h6>
  <div class="star-rating">
      ${[1, 2, 3, 4, 5].map(star => `
        <span class="star" data-value="${star}">&#9733;</span>
      `).join('')}
    </div>
      <textarea class="form-control mt-2" placeholder="Write your review here..."></textarea>
      <button class="btn btn-secondary mt-2">Submit Review</button>
    </div>
    <div class="favourite-section mt-3">
      <button class="btn btn-warning" id="favouriteButton">
        <i class="fas fa-star"></i> Favourite
      </button>
    </div>
  </div>
</div>

`;
    // Add event listener for the favourite button
    document.getElementById("favouriteButton").addEventListener("click", function () {
        attractionData.recommended_count += 1;
        // Optionally, update the UI to reflect the new recommended count
        console.log(`Attraction recommended count: ${attractionData.recommended_count}`);
        // Add to user's favourites (this would typically involve an API call)
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
            // Optionally, submit the rating (this would typically involve an API call)
        });
    });
}

function loadCarouselImages() {
    const images = [
        { src: "images/image1.jpg", alt: "First slide" },
        { src: "images/image2.jpg", alt: "Second slide" },
        { src: "images/image3.jpg", alt: "Third slide" }
    ];
    return images.map((img, index) => `
        <div class="carousel-item ${index === 0 ? 'active' : ''}">
            <img src="${img.src}" class="d-block object-fit-cover carousel-images w-100" alt="${img.alt}">
        </div>
    `).join('');
}
function hideSidebarContent() {
    const selectedAttractionsInfo = document.getElementById("selectedAttractionsInformation");
    const noAttractionSelected = document.getElementById("NoAttractionSelected");
    selectedAttractionsInfo.style.display = "none";
    noAttractionSelected.style.display = "block";
}
