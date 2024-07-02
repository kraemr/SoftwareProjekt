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
            <span class="user-settings-button w-100" onclick="OpenLogout()">Sign out</span>
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
    var apiUrl = document.location.origin + '/api/faourites';
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            console.log(data);
            document.getElementById('sidebar-content').innerHTML = `
        <div class="col-12">
            <h4 class="mb-3">Favourites</h4>
        </div>
        ${data.map(favourite => `
            <div class="col-12">
                <div class="card mb-3">
                    <div class="card-body">
                        <h5 class="card-title">${favourite.title}</h5>
                        <p class="card-text">${favourite.info}</p>
                        <p class="card-text">${favourite.city}</p>
                        <p class="card-text">${favourite.type}</p>
                        <p class="card-text">${favourite.Stars}</p>
                        <p class="card-text">${favourite.recommended_count}</p>
                    </div>
                </div>
            </div>
        `).join('')}
    `;
        }
        );
}

