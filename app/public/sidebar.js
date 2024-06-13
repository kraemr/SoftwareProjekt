// Toggle user settings
function toggleSettings() {
    var settings = document.getElementById('toggle-settings');
    if (settings.style.display === 'none') {
        settings.style.display = 'block';
    } else {
        settings.style.display = 'none';
    }
}
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