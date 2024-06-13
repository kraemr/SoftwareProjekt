function toggleSettings() {
    const settingsContainer = document.querySelector('.settings-container');
    settingsContainer.style.display = (settingsContainer.style.display === 'block') ? 'none' : 'block';
  }
  
  // Optional: Close the settings when clicking outside of it
  document.addEventListener('click', function(event) {
    const userContainer = document.querySelector('.user-container');
    const settingsContainer = document.querySelector('.settings-container');
    if (!userContainer.contains(event.target) && !settingsContainer.contains(event.target)) {
      settingsContainer.style.display = 'none';
    }
  });
  
// Toggle left sidebar
function toggleNav() {
    const sidepanel = document.getElementById("mySidepanel");
    const button = document.getElementById("toggleButton");
    if (window.innerWidth <= 768) { // Check if the screen size is small
        button.style.left = "50%";
        sidepanel.style.width = "100%";
        if (sidepanel.style.height === "40%") {
            sidepanel.style.height = "0";
            button.style.bottom = "2px";
        } else {
            sidepanel.style.height = "40%";
            button.style.bottom = "41%";
        }
    } else {
        button.style.left = "2px";
        if (sidepanel.style.width === "25%") {
            sidepanel.style.width = "0";
            button.style.left = "2px";
        } else {
            sidepanel.style.width = "25%";
            button.style.left = "20%";
        }
    }
}