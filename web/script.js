// About button functionality
const aboutButton = document.getElementById('about');
const aboutModal = document.getElementById('aboutModal');
const closeAboutModal = document.getElementsByClassName('close')[0];

const loginModal = document.getElementById('loginModal');
const closeLoginModal = document.getElementsByClassName('close')[1]; // Close button for Login modal
const mainPage = document.getElementById('mainPage');

// Open the modal when the About button is clicked
aboutButton.addEventListener('click', () => {
    aboutModal.style.display = 'block';
});

// Close the modal when the "x" is clicked
closeAboutModal.addEventListener('click', () => {
    aboutModal.style.display = 'none';
});

// Close the modal when the user clicks anywhere outside of the modal
window.addEventListener('click', (event) => {
    if (event.target === aboutModal) {
        aboutModal.style.display = 'none';
    }
});

// Function to show the login modal
function showLoginModal() {
  loginModal.style.display = 'block';
  mainPage.classList.add('blur'); // Apply blur to the background
}

// Document ready event
document.addEventListener('DOMContentLoaded', (event) => {
  // Show the login modal when the document is fully loaded
  showLoginModal();
});
