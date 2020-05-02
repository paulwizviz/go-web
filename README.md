![build](https://github.com/paulwizviz/go-web/workflows/build/badge.svg)

# Overview

This a template to help you scaffold a project to create macOS, Linux and Windows native app or Docker base app incorporating a Web frontend (e.g. ReactJS, Vue, etc) into a single package.

This project is **NOT** a Go module or library where you would `go get` packages for inclusion into your project.

The scaffold generated from this template includes a sample with a frontend implemented using ReactJS. The sample is intended to demonstrates steps involve in building functional apps. The sample does **NOT** represent any specific use case but you can extend it into a fully functional production native or container app. You may also elect to replace the sample app with one configured to work within the context of the scaffold or to modify and extend the sample to meet your requirement.

## How do I scaffold a project from this template?

At the top this page, you should see a green coloured button named `Use this template`. Click on the button and follow the instructions from Github.

## What do I get out-of-the-box when I create a scaffold?

You'll get the following:

* A ReactJS sub-project.

* Go codes to manage app configuration, startup sequence and network operations.

* Go implementation of a webserver and RESTful server

* Build scripts -- Docker based -- to create native (macOS, Linux and Windows) app or Docker image app.

* A locally deployable development environment

## Prerequisite

In order to build apps based on this template ensure you have the following items installed:

* Docker and Docker compose

* [Go toolkit version 1.11 onwards](https://blog.golang.org/)

* Node

## How do I see the sample in action?

There are two options for you to see the sample in action. In both cases, you will need to build the sample first. Please follow this [doc](./docs/Build.md) for instruction.

### Native apps

To see a native app in action, follow these steps:

1. Open an terminal and navigate to `./build/native/<platform of your choice>`

2. You have the following options:

    2.1. To run the app with a UI frontend, execute the command `goreact start ui` (for macOS and Linux) or `goreact.exe start ui` (for Windows).

    2.2. To run the app with no UI frontend (i.e. headless), execute the command `goreact start noui` (for macOS and Linux) or `goreact.exec start noui`

NOTE: By default the app will run require port 80 to be available. If you need the app to use other ports type the command `goreact --help`.

3. Assuming you have no problem with STEP 2.1 and operating on port 80, open a browser with url to `localhost`.

### Container app

1. Execute the following command `./scripts/test/e2e.sh start`. This will start-up a container with port 80 open.

NOTE: If you wish to have the container expose a different port modify the docker-compose deployment found here [script](./deployments/e2e/docker-compose.yaml)

2. Assuming you have no problem with STEP 1 and operating on port 80, open a browser with url to `localhost`.

3. To stop you app, execute the following command `./scripts/test/e2e.sh stop`.

4. To clean up your workspace, `./scripts/test/e2e.sh clean`.

## How do I extend/modify my scaffold?

Please refer to:

* [Architecture](./docs/Arch.md)

* [Build](./docs/Build.md)
 
