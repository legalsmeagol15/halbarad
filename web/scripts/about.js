// About button functionality
const aboutButton = document.getElementById('about');
const aboutModal = document.getElementById('aboutModal');
const closeAboutModal = document.getElementsByClassName('close')[0];

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
