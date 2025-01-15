const loginModal = document.getElementById('loginModal');
const mainPage = document.getElementById('mainPage');
const switchToRegisterButton = document.getElementById('switchToRegister');
const switchToLoginButton = document.getElementById('switchToLogin');
const loginForm = document.getElementById('loginForm');
const registerForm = document.getElementById('registerForm');


// Function to show the login modal
function showLoginModal() {
    loginModal.style.display = 'block';
    mainPage.classList.add('blur'); // Apply blur to the background
    loginForm.style.display = "block";
    registerForm.style.display = "none";
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
  