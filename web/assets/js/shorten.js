const URL_INPUT = document.getElementById("urlInput");
const SUBMIT_BUTTON = document.getElementById("submitButton");
const URL_TABLE = document.getElementById("urlTable")

// Add an event listener for clicking the submit button
SUBMIT_BUTTON.addEventListener("click", () => {
    // Check URL input validity
    if (!URL_INPUT.checkValidity()) return;

    const SHORT_PATH = generateRandomString();
    const LONG_URL = URL_INPUT.value;
    const RESPONSE = callShortEndpoint(SHORT_PATH, LONG_URL);

    RESPONSE.then(() => {
        // New table row
        const NEW_ROW = URL_TABLE.insertRow();
        let newCell = NEW_ROW.insertCell();
        newCell.appendChild(document.createTextNode(location.origin + "/" + SHORT_PATH));

        newCell = NEW_ROW.insertCell();
        newCell.appendChild(document.createTextNode(LONG_URL));

        // Change table visibility
        URL_TABLE.removeAttribute("hidden");
    });
});

// Make a request on the shortening API endpoint
async function callShortEndpoint(shortPath, longUrl) {
    return await fetch(location.origin + "/api/shorten", {
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
    const CHARACTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    let randomString = "";

    for (let i = 0; i < 6; i++) {
        randomString += CHARACTERS.charAt(Math.floor(Math.random() * CHARACTERS.length));
    }

    return randomString;
}
