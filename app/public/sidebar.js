
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
function toggleNav() {
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
            button.style.left = "20%";
        }
    }
}