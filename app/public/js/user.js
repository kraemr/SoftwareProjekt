var userID;
var username;
var userImage;

// Unused function to get the user information by their email
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
// Unused function to set the user information
function setUserSettings() {
  const settingsContainer = document.querySelector(".user-icon-src");
  settingsContainer.style.backgroundImage = "url('/images/user-icon.png')";
}
// Function that returns true of false when asked if the user is logged in
function isLoggedIn() {
  return fetch(document.location.origin + "/api/logged_in", {
    method: "GET",
  })
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      if (data.success == true) {
        return true;
      } else {
        return false;
      }
    })
    .catch((error) => {
      console.error("Error checking login status:", error);
      return false;
    });
}
// Unused function to log out the current user
function logoutCurrentUser() {
  fetch(document.location.origin + "/api/logout", {
    method: "GET",
  });
}
