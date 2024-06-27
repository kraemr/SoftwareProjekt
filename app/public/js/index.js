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
            <span class="user-settings-button w-100" onclick="OpenHistory()">History</span>
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

