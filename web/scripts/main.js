import * as THREE from 'three';

let scale = 1.0;
const rescaleFactor = 1.1;
let offsetX = 0, offsetY = 0;

function getMouse(element, event){
    let rect = element.getBoundingClientRect();
    const mouseX = ((event.clientX - rect.left) / scale) + offsetX;
    const mouseY = ((event.clientY - rect.top) / scale) + offsetY;
    return [mouseX, mouseY];
}

// Current view parameters display
const coordsDisplay = document.getElementById('current-coords');

// Button event listeners
document.getElementById('view-front').addEventListener('click', () => {
    // set view to front orientation
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


// THE CANVAS ITSELF
const canvas = document.getElementById('main-canvas');

canvas.addEventListener('mousemove', (event) => {
    const [mouseX, mouseY] = getMouse(canvas, event);
    coordsDisplay.textContent = "x:" + mouseX + ", y:" + mouseY;
});


canvas.addEventListener('wheel', (event) => {
    event.preventDefault();
    const [mouseX, mouseY] = getMouse(canvas, event);
    
    let newScale = scale;
    if (event.deltaY < 0) {
        newScale *= rescaleFactor;
    } else {
        newScale /= rescaleFactor;
    }
    offsetX = mouseX - ((mouseX - offsetX) * (newScale / scale));
    offsetY = mouseY - ((mouseY - offsetY) * (newScale / scale));
    scale = newScale;

    console.log('Scale:', scale, "  mouseX:", mouseX, "  mouseY:", mouseY, "  offsetX:", offsetX, "  offsetY:", offsetY);
    draw();
});

const container = document.getElementById('canvas-container');
container.addEventListener('scroll', () => {
    draw(); // redraw with new offset
});

function draw() {
    
    const ctx = canvas.getContext('2d');
    ctx.imageSmoothingEnabled = false;
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    // None of these work!
    // ctx['imageSmoothingEnabled'] = false;       /* standard */
    // ctx['mozImageSmoothingEnabled'] = false;    /* Firefox */
    // ctx['oImageSmoothingEnabled'] = false;      /* Opera */
    // ctx['webkitImageSmoothingEnabled'] = false; /* Safari */
    // ctx['msImageSmoothingEnabled'] = false;     /* IE */
    
    // SOME EXAMPLE LINES FOR TESTING
    ctx.lineCap = 'butt';
    ctx.lineJoin = 'miter';
    ctx.save();
    ctx.setTransform(scale, 0, 0, scale, offsetX, offsetY);

    ctx.strokeStyle = 'blue';
    ctx.lineWidth = 3;
    ctx.beginPath();
    ctx.moveTo(5, 25);
    ctx.lineTo(300, 25);
    ctx.stroke();

    // ANOTHER EXAMPLE LINE FOR TESTING
    ctx.strokeStyle = 'green';
    ctx.lineWidth = 2;
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

    ctx.restore();
}
