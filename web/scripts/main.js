//window.location.href = 'main.html';

// Example: update current view parameters display
function updateCurrentViewParams(paramsText) {
    document.getElementById('current-view-params').textContent = 'Current view: ' + paramsText;
}

// Button event listeners
document.getElementById('view-front').addEventListener('click', () => {
    // set view to front orientation
});
document.getElementById('view-top').addEventListener('click', () => {
    // set view to top orientation
});
document.getElementById('view-side').addEventListener('click', () => {
    // set view to side orientation
});

// Entry box event (e.g., on enter)
document.getElementById('view-input').addEventListener('keydown', (event) => {
    if (event.key === 'Enter') {
        const params = event.target.value;
        // apply view param logic here
        updateCurrentViewParams(params);
    }
});



// VIEW CONTROL
const canvasContainer = document.getElementById('canvas-container');
window.addEventListener('DOMContentLoaded', () => {
    draw(); // draw initial scene
});


canvasContainer.addEventListener('scroll', () => {
    const scrollX = canvasContainer.scrollLeft;
    const scrollY = canvasContainer.scrollTop;
    console.log('Scroll:', scrollX, scrollY);
});

const canvas = document.getElementById('main-canvas');


const container = document.getElementById('canvas-container');
container.addEventListener('scroll', () => {
    draw(); // redraw with new offset
});

function draw() {
    
    const ctx = canvas.getContext('2d');
    ctx['imageSmoothingEnabled'] = false;
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    
    // SOME EXAMPLE LINES FOR TESTING
    ctx.lineCap = 'butt';
    ctx.lineJoin = 'miter';

    ctx.strokeStyle = 'blue';
    ctx.lineWidth = 3;
    ctx.beginPath();
    ctx.moveTo(5, 25);
    ctx.lineTo(300, 25);
    ctx.stroke();

    // ANOTHER EXAMPLE LINE FOR TESTING
    ctx.strokeStyle = 'green';
    ctx.beginPath();
    ctx.moveTo(10.5, 30.5);
    ctx.lineTo(305.5, 35.5);
    ctx.stroke();

    // WHAT WOULD A BEZIER CURVE LOOK LIKE?
    ctx.strokeStyle = 'red';
    ctx.beginPath();
    ctx.moveTo(0, 5);
    ctx.bezierCurveTo(50, 20, 25, 10, 35, 15);
    ctx.stroke();
}
