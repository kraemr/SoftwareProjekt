// Function to display the favourites in the sidebar
function OpenFavourites() {
    hideSidebarContent();
    openSidepanel();
    // Get the favourites from the current user
    fetch(document.location.origin + "/api/favorites", {
        method: "GET",
    })
        .then(response => response.json())
        .then(data => {
            const sidebarContent = document.getElementById("showFavourites");
            sidebarContent.style.display = "block";
            // Filter input
            sidebarContent.innerHTML = `
                <div class="input-group mb-3">
                    <input type="text" class="form-control" id="filter-input" placeholder="Filter favourites" oninput="filterFavourites()">
                </div>
            `;
            console.log(data);
            // Display the favourites in cards
            data.forEach(favorite => {
                const card = document.createElement("div");
                card.className = "card mb-2 m-2";

                const cardBody = document.createElement("div");
                cardBody.className = "card-body";

                const cardTitle = document.createElement("h5");
                cardTitle.className = "card-title";
                cardTitle.innerHTML = favorite.title;

                const cardText = document.createElement("p");
                cardText.className = "card-text";
                cardText.innerHTML = favorite.info;

                const buttonContainer = document.createElement("div");
                buttonContainer.style.display = "flex";
                buttonContainer.style.justifyContent = "space-between";

                // Add a button to delete the favourite
                const deleteButton = document.createElement("button");
                deleteButton.className = "btn btn-danger";
                deleteButton.innerHTML = "Unfavourite";
                deleteButton.onclick = function() {
                     // use the attraction ID from the favorite to delete the attraction data
                    fetch(document.location.origin + "/api/favorites?id=" + favorite.Id, {
                        method: "DELETE",
                    })
                    .then(response => response.json())
                    .then(result => {
                        if (result.success) {
                            console.log('Favorite deleted:', result);
                            card.remove();
                        } else {
                            console.error('Error deleting favorite:', result.info);
                        }
                    })
                    .catch(error => {
                        console.error('Error deleting favorite:', error);
                    });
                };
                // Add a button to show the favourite on the map
                const showOnMapButton = document.createElement("button");
                showOnMapButton.className = "btn btn-primary";
                showOnMapButton.innerHTML = "Show on Map";
                showOnMapButton.onclick = function() {
                    // use the attraction ID from the favorite to fetch the attraction data
                    fetch(document.location.origin + "/api/attractions?id=" + favorite.Attraction_id, {
                        method: "GET",
                    })
                    .then(response => response.json())
                    .then(data => {
                        console.log('Attraction data:', data);
                        showAttractionOnMap(data[0]);
                    })
                };

                cardBody.appendChild(cardTitle);
                cardBody.appendChild(cardText);
                card.appendChild(cardBody);
                sidebarContent.appendChild(card);
                cardBody.appendChild(buttonContainer);
                buttonContainer.appendChild(deleteButton);
                buttonContainer.appendChild(showOnMapButton);
                console.log(data);
            });
        })
        .catch(error => {
            console.error('Error fetching favorites:', error);
        });
}
// Function to filter the favourites by anything in the card body
function filterFavourites() {
    const filterValue = document.getElementById('filter-input').value.toLowerCase();
    const favouriteCards = document.querySelectorAll('#showFavourites .card');

    favouriteCards.forEach(card => {
        const cardText = card.querySelector('.card-body').innerText.toLowerCase();
        if (cardText.includes(filterValue)) {
            card.style.display = '';
        } else {
            card.style.display = 'none';
        }
    });
}