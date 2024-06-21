
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
                }
            });
        });

    }
}
function loadSidebarInfo() {
    const selectedAttractionsInfo = document.getElementById("selectedAttractionsInformation");
    const noAttractionSelected = document.getElementById("NoAttractionSelected");
    if (selectedAttractionsInfo.style.display === "") {
        console.log("No attraction selected");
        noAttractionSelected.style.display = "block";
        noAttractionSelected.innerHTML = `
<div class="card bg-black m-2">
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
    <h5 class="card-title">Attraction Name</h5>
    <p class="card-text">Some quick example text to build on the attraction name and make up the bulk of the card's content.</p>
    <a href="#" class="btn btn-primary">Go somewhere</a>
  </div>
</div>

`;
    } else {
        noAttractionSelected.style.display = "none";
        selectedAttractionsInfo.style.display = "block";
        selectedAttractionsInfo.innerHTML = `
    <div>
        Test2
    </div>
`;
    }
}

function loadCarouselImages() {
    const images = [
        { src: "path_to_first_image.jpg", alt: "First slide" },
        { src: "path_to_second_image.jpg", alt: "Second slide" },
        { src: "path_to_third_image.jpg", alt: "Third slide" }
    ];
    return images.map((img, index) => `
        <div class="carousel-item ${index === 0 ? 'active' : ''}">
            <img src="${img.src}" class="d-block w-100" alt="${img.alt}">
        </div>
    `).join('');
}
function loadMarkerInfoToSidebar(attractionData) {
    const sidebar = document.getElementById("selectedAttractionsInformation");
    sidebar.innerHTML = attractionData.name;
    console.log("success clicking loadMarkerInfoToSidebar");
}
