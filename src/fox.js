(async () => {
	console.log("Connecting to native host...");
	const host = await browser.runtime.connectNative("foxpi");

    console.log("Connected to native host.");

    host.onDisconnect.addListener(() => {
        console.log("Disconnected from native host.");
    });

	host.onMessage.addListener(async url => {
        if (!url || !url.startsWith("http://") && !url.startsWith("https://")) {
            console.log(`Unable to fetch invalid url: ${url}`);

            return;
        }

        console.log(`Fetching "${url}"...`);

        const response = await read(url);

        console.log(`Sending response: ${response?.length || 0} bytes`);

        host.postMessage(response);
	});
})();

export async function read(url) {
    try {
        const response = await fetch(url);

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
