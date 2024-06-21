var userID;
var username;
var userImage;
function getUserByEmail(email) {
    fetch(document.location.origin + "/api/users?email=" + email, {
        method: "GET",
    })
        .then((response) => {
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            return response.json();
        })
        .then((data) => {
            console.log(data);
            userID = data[0].id;
        })
        .catch((error) => {
            console.error("There was a problem with the request:", error);
            alert("Failed to load markers. Please try again.");
        });
}
function setUserSettings() {
    const settingsContainer = document.querySelector('.user-icon-src');
    settingsContainer.style.backgroundImage = "url('/images/user-icon.png')";
}
function isLoggedIn(){
    return fetch(document.location.origin + "/api/logged_in", {
        method: "GET",
    })
    .then((response) => {
        if (!response.ok) {
            throw new Error("Network response was not ok");
        }
        return response.status === 200;
    })
    .catch((error) => {
        console.error("Error checking login status:", error);
        return false;
    });
}
function logoutCurrentUser() {
    fetch(document.location.origin + "/api/logout", {
        method: "GET",
    });
}