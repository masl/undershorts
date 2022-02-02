const URL_INPUT = document.getElementById("urlInput");
const SUBMIT_BUTTON = document.getElementById("submitButton");
const SHORTENED_PARAGRAPH = document.getElementById("shortened");

const SHORT_PATH = generateRandomString();

// Add an event listener for clicking the submit button
SUBMIT_BUTTON.addEventListener("click", () => {
    const LONG_URL = URL_INPUT.value;
    const response = callShortEndpoint(SHORT_PATH, LONG_URL);

    response.then(() => {
        alert("URL shortened!");
        SHORTENED_PARAGRAPH.innerHTML = location.origin + "/" + SHORT_PATH;
    });
});

// Make a request on the shortening API endpoint
function callShortEndpoint(shortPath, longUrl) {
    return fetch(location.origin + "/api/shorten", {
        //mode: "no-cors",
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            "shortPath": shortPath,
            "longUrl": longUrl
        })
    });
}

// Generate a random 6 letter long string
function generateRandomString() {
    const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    let string = "";

    for (let i = 0; i < 6; i++) {
        string += characters.charAt(Math.floor(Math.random() * characters.length));
    }
    return string;
}

//TODO: Adjust js to new html layout