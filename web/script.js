// About button functionality
const aboutButton = document.getElementById('about');
const aboutModal = document.getElementById('aboutModal');
const closeAboutModal = document.getElementsByClassName('close')[0];

const loginModal = document.getElementById('loginModal');
const closeLoginModal = document.getElementsByClassName('close')[1]; // Close button for Login modal
const mainPage = document.getElementById('mainPage');
const switchToRegisterButton = document.getElementById('switchToRegister');
const switchToLoginButton = document.getElementById('switchToLogin');
const loginForm = document.getElementById('loginForm');
const registerForm = document.getElementById('registerForm');
const modalTitle = document.getElementById('modalTitle');

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

// Switch to Register Form
switchToRegisterButton.addEventListener('click', () => {
  loginForm.style.display = 'none';
  registerForm.style.display = 'block';
  //modalTitle.textContent = 'Register';
});

// Switch to Login Form
switchToLoginButton.addEventListener('click', () => {
  registerForm.style.display = 'none';
  loginForm.style.display = 'block';
  //modalTitle.textContent = 'Login';
});

// Handle Login Form submission
document.getElementById('loginForm').addEventListener('submit', function(event) {
  console.log("Click")
  event.preventDefault(); // Prevent the default form submission

  // Get the values from the input fields
  const username = document.getElementById('username').value;
  const password = document.getElementById('password').value;

  // Basic validation
  if (username === '' || password === '') {
      document.getElementById('message').innerText = 'Please fill in all fields.';
      return;
  }
});

// Handle register Form submission
document.getElementById('registerForm').addEventListener('submit', function(event) {
  event.preventDefault(); // Prevent the default form submission

  // Get the values from the input fields
  const username = document.getElementById('username').value;
  const password = document.getElementById('password').value;

  // Basic validation
  if (username === '' || password === '') {
      document.getElementById('message').innerText = 'Please fill in all fields.';
      return;
  }
});
