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

    if (sidepanel.style.width === "20%") {
      sidepanel.style.width = "0";
      button.style.left = "2px";
    } else {
      sidepanel.style.width = "20%";
      button.style.left = "20%";
    }
  }