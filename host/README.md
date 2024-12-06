This is the [native messaging](https://developer.mozilla.org/en-US/docs/Web/API/Native_Messaging_API) host for [foxpi](https://github.com/coalaura/foxpi). It provides the http server for requests and handles communication with the extension.

### Requirements
- [golang](https://golang.org)

### Installation
1. Download the host source and run the build script.
2. The build script will create a directory called `.foxpi` in your user profile.
3. It will also try to add a registry key, if you don't run the script as administrator you need to add it manually.
4. Once properly installed, you're good to install the extension itself.