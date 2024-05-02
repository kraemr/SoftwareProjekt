// Get the sidebar element
const sidebar = document.querySelector('.sidebar');

// Get the toggle button element
const toggleButton = document.querySelector('.toggle-sidebar-button');

// Add event listener to the toggle button
toggleButton.addEventListener('click', () => {
    // Hide the sidebar
    sidebar.classList.toggle('hide');
});