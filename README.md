![build](https://github.com/paulwizviz/go-web/workflows/build/badge.svg)

# Overview

This a template to help you scaffold a project for creating macOS, Linux and Windows native app or Docker base app incorporating a Web-based frontend (e.g. ReactJS, Vue, etc) with a web and RESTFul server. 

To understand how the app is built please refer to this [doc](./docs/Build.md).

**NOTE:** This project is **NOT** a Go module or library where you would `go get` packages for inclusion into your project. Neither is it an npm module.

## How do I scaffold a project from this template?

At the top this page, you should see a green coloured button named `Use this template`. Click on the button and follow the instructions from Github.

## What do I get out-of-the-box when I create a scaffold?

You'll get the following:

* A ReactJS-based skeleton sub-project.

* Go codes to create commandlines operation to configuration, startup sequence and network operations.

* Go based skeleton webserver and RESTful server.

* Docker based build scripts to help you create native (macOS, Linux and Windows) app or Docker image app.

* A locally deployable development environment.

## Prerequisite

In order to build apps based on this template, you **must** install Docker and Docker compose.

You should also install the following items to help you edit source codes:

* [Go toolkit version 1.11 onwards](https://blog.golang.org/)

* Node

NOTE: These toolkits are not use to build your projects.

## How do I extend/modify my scaffold?

Please refer to the [architecture docs](./docs/Arch.md)

## How do I build my project?

Please refer to the [build doc](./docs/Build.md).

## How do I see my projects in action?

Please refer to [running the app doc](./docs/Run.md).

 
