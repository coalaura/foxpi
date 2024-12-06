# FoxPI
Control firefox remotely to perform requests and bypass bot detection.

### Requirements
- Firefox version 133 or higher
- Native messaging host [foxpi](host)

### Installation
1. Download the native messaging host and run the build script.
2. The build script will create a directory called `.foxpi` in your user profile.
3. It will also try to add a registry key, if you don't run the script as administrator you need to add it manually.
4. Install the extension.
5. You're good to go! You can perform requests using the local http server: `http://localhost:4269/`

### Usage
You can perform requests using the local http server: `http://localhost:4269/`

|Query parameter|Description|
|---|---|
|`method`|The HTTP method to use (defaults to `GET`).|
|`url`|The URL to request (e.g. `https://example.com`).|

**Example:**
```
http://localhost:4269/?method=GET&url=https://google.com
```

The response will be the response from the request if it was successful. The forwarded response includes the headers and the body of the response.