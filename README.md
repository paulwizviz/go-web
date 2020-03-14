# Overview

This a template to help you creating a native macOS, Linux, Windows and/or Container app powered by a combination of Web UI frameworks (ReactJS, Vue, etc) and Go framework. 

This project is **NOT** a Go module or library where you would `go get` packages fore inclusion into your project. You will need to create a scaffold to use it. The scafold is presented in a way to help you extend, modify or create new features.

There is a sample app is primarily to demonstrates steps involve in building a functional apps. The sample app does **NOT** represent any specific use case but you can use it as a basis to build a fully functional production release or deployment. Or you can replace the sample app with equivalent implementation and project configurations.

## How do I scaffold a project from this template?

Click on the template button and follow the instructions from Github.

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

## How do I see the sample app in action?

1. Assuming you have already clone this project, click on this [doc](./docs/Build.md) and follow the steps to build a representative production native app. 

2. Open an terminal and navigate to `./build/package/<platform of your choice>`.

3. Activate the app `go-react` (for macOS and Linux) or `go-react.exe` (for windows). By default the app will run require port 80 to be available. If you need the app to use other ports type the command `go-react --help`.

4. Assuming you have no problem with STEP 3, open a browser with url to `localhost`.

## How do I extend/modify my scaffold?

Please refer to:

* [Architecture](./docs/Arch.md)

* [Build](./docs/Build.md)
 
