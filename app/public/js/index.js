function loadUserSettings() {
    isLoggedIn().then(isUserLoggedIn => {
        console.log(isUserLoggedIn);
        if (isUserLoggedIn) {
            document.getElementById('user-settings').innerHTML = `
        <div class="col-12">
            <span class="user-settings-button w-100" onclick="OpenProfile()">Your  Profile</span>
        </div>
        <div class="col-12">
            <span class="user-settings-button w-100" onclick="OpenFavourites()">Favourites</span>
        </div>
        <div class="col-12">
            <span class="user-settings-button w-100" onclick="logoutBtnClick()">Sign out</span>
        </div>
    `;
        } else {
            document.getElementById('user-settings').innerHTML = `
        <div class="col-12 mb-2">
            <a href="login.html" class="btn btn-primary w-100">Login</a>
        </div>
        <div class="col-12">
            <a href="register.html" class="btn btn-primary w-100">Register</a>
        </div>
    `;
        }
    });
}

function OpenProfile() {
    window.location.href = "your_Profile.html";
}
function OpenFavourites() {
    hideSidebarContent();
    var apiUrl = document.location.origin + '/api/favorites';
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            const sidebarContent = document.getElementById("showFavourites");
            console.log(data);
            data.forEach(favorite => {
                const card = document.createElement("div");
                card.className = "card mb-2";
                
                const cardBody = document.createElement("div");
                cardBody.className = "card-body";
                
                const cardTitle = document.createElement("h5");
                cardTitle.className = "card-title";
                cardTitle.innerHTML = favorite.name;
                
                const cardText = document.createElement("p");
                cardText.className = "card-text";
                cardText.innerHTML = favorite.description;
                
                cardBody.appendChild(cardTitle);
                cardBody.appendChild(cardText);
                card.appendChild(cardBody);
                sidebarContent.appendChild(card);
            });
        })
        .catch(error => {
            console.error('Error fetching favorites:', error);
        });
}


