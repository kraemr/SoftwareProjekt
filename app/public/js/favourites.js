function OpenFavourites() {
    hideSidebarContent();
    fetch(document.location.origin + "/api/attractions", {
        method: "GET",
    })
        .then(response => response.json())
        .then(data => {
            const sidebarContent = document.getElementById("showFavourites");
            sidebarContent.style.display = "block";
            console.log(data);
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

                const deleteButton = document.createElement("button");
                deleteButton.className = "btn btn-danger";
                deleteButton.innerHTML = "Delete";
                deleteButton.onclick = function() {
                    fetch(document.location.origin + "/api/favorites", {
                        method: "DELETE",
                    })
                    .then(response => response.json())
                    .then(result => {
                        if (result.success) {
                            card.remove();
                        } else {
                            console.error('Error deleting favorite:', result.info);
                        }
                    })
                    .catch(error => {
                        console.error('Error deleting favorite:', error);
                    });
                };

                const showOnMapButton = document.createElement("button");
                showOnMapButton.className = "btn btn-primary";
                showOnMapButton.innerHTML = "Show on Map";
                showOnMapButton.onclick = function() {
                    showAttractionOnMap(favorite);
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