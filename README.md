# Overview

This a template to help you scaffold a project for creating a native macOS, Linux, Windows and/or Container app powered by a combination of Web UI frameworks (ReactJS, Vue, etc) and Go framework. 

This is **NOT** a Go module or library.

## What do I get out-of-the-box when I create a scaffold?

You'll get the following:

1. An example ReactJS frontend that you can extend or use it to learn concept and then replace with other equivalent framework such as Vue.

2. An environment to support both Web and Go development.

3. Build scripts -- Docker based -- to create native (macOS, Linux and Windows) app or container based app.

4. The underlying codes are based on Go modules (i.e. Go11+).

## Prerequisite

In order to build apps based on this template ensure you have the following items installed:

1. Docker and Docker compose

2. [Go toolkit version 1.11 onwards](https://blog.golang.org/)

3. (Optional) Node (out-of-the-box you don't need this because this template uses Docker to build the node part)

## How to scaffold a project from this template?

Click on the template button and follow the instructions from Github.

## Is there a sample app I can use?

Yes follow these steps:

1. Click on this [doc](./docs/Build.md) and build a production native app. 

2. Open an terminal and navigate to `./build/package/<platform of your choice>`.

3. Activate the app `go-react` (for macOS and Linux) or `go-react.exe` (for windows). By default the app will run require port 80 to be available. If you need the app to use other ports type the command `go-react --help`.

4. Assuming you have no problem with STEP 3, open a browser with url to `localhost`.

NOTE: The sample app is only intented to illustrate concept and if appropriate for you to extend into a full production or mission ready application. The sample app is **NOT** intended for production use.

## How do I extend/modify my scaffold?

Please refer to:

* [Architecture](./docs/Arch.md)

* [Build](./docs/Build.md)
 
