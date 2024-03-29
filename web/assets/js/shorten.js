const URL_INPUT = document.getElementById("urlInput");
const SUBMIT_BUTTON = document.getElementById("submitButton");
const URL_TABLE = document.getElementById("urlTable")
const NOTIFICATION = document.getElementById("alert")

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
        const SHORT_URL = location.origin + "/" + SHORT_PATH;
        let newCell = NEW_ROW.insertCell();
        newCell.appendChild(document.createTextNode(SHORT_URL));

        // Copy on click
        newCell.style.cursor = "copy";
        newCell.addEventListener("click", () => {
            navigator.clipboard.writeText(SHORT_URL).then(() => {
                showNotification()
            }, reason => {
                console.log("Rejected:" + reason);
            });
        })

        newCell = NEW_ROW.insertCell();
        newCell.appendChild(document.createTextNode(LONG_URL));

        // Change table visibility
        URL_TABLE.removeAttribute("hidden");
    });
});

// Make a request on the shortening API endpoint
async function callShortEndpoint(shortPath, longUrl) {
    return await fetch(location.origin + "/api/v1/shorten", {
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

function showNotification() {

    NOTIFICATION.classList.remove("alertFadeout");
    NOTIFICATION.classList.add("alertFadein");

    setTimeout(() => {
        NOTIFICATION.classList.remove("alertFadein");
        NOTIFICATION.classList.add("alertFadeout");
        //NOTIFICATION.style.display = "none";
    }, 2000)
}
