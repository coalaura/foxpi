const ValidMethods = [
    "GET", "POST", "PUT", "PATCH", "DELETE"
];

(async () => {
	console.log("Connecting to native host...");
	const host = await browser.runtime.connectNative("foxpi");

    console.log("Connected to native host.");

    host.onDisconnect.addListener(() => {
        console.log("Disconnected from native host.");
    });

	host.onMessage.addListener(async request => {
        const validated = validateRequest(request);

        if (!validated) {
            console.log("Invalid request.", request);

            return;
        }

        console.log(`Fetching "${request.method} ${request.url}"...`);

        const response = await read(request);

        console.log(`Sending response: ${response?.length || 0} bytes`);

        host.postMessage(response);
	});
})();

function validateRequest(request) {
    let { url, method, headers } = request;

    // Default values for method and headers if not provided
    if (!method) method = "GET";
    if (!headers) headers = {};

    // Ensure URL is valid
    if (!url || (!url.startsWith("http://") && !url.startsWith("https://"))) {
        return false;
    }

    // Ensure method is valid
    if (!ValidMethods.includes(method)) {
        return false;
    }

    // Ensure headers is an object
    if (typeof headers !== "object") {
        return false;
    }

    // Ensure headers is a valid object
    for (const key in headers) {
        if (typeof key !== "string") {
            return false;
        }

        const value = headers[key];

        if (typeof value !== "string") {
            return false;
        }
    }

    return {
        url: url,
        method: method,
        headers: headers
    };
}

export async function read(request) {
    try {
        const response = await fetch(request.url, {
            method: request.method,
            headers: request.headers
        });

        if (!response.ok) {
            return false;
        }

        const buffer = await response.arrayBuffer(),
            bytes = new Uint8Array(buffer);

        return {
            headers: response.headers,
            body: btoa(String.fromCharCode(...bytes)),
            length: bytes.length
        };
    } catch {}

    return false;
}
