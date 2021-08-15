![build](https://github.com/paulwizviz/go-web/workflows/build/badge.svg)

# Overview

This is a template to help developers scaffold a project for creating macOS, Linux and Windows native app or Docker base app incorporating a Web-based frontend (ReactJS, Vue, etc) with Go bases backend services (Web server, REST server and etc).

What problem is this template meant to solve?

Often Go developers are faced with having to provide a user interface to support backend services. For Go developers that do not want spend time writing frontend codes, they would like to be able to use pre-developed web interface and integrate as part of their native app or web-based project.

## Prerequisite

You **must** install Docker and Docker compose as these toolkits are needed to build native apps and/or container-based app from the scaffold plus your extensions.

You will also need to install these artefacts:

* [Go toolkit version 1.16 onward](https://blog.golang.org/);

* Node (version depends on the web framework you are using).

## <a name="out-of-the-box">What do I get out-of-the-box when I create a scaffold</a>?

When you derive a project directly from this template, you will get the following:

* A layout based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout);

* A ReactJS-based skeleton sub-project;

* A skeleton Go source code to generate webserver and REST server;

* Docker based build scripts to help you create native (macOS, Linux and Windows) app or Docker image app.

* A locally deployable development environment to help you see the interactions between your web source code and Go REST API (including hot code reloading whilst you develop);

* Docker-based testing framework. 

## How do I use this template?

Out-of-the-box there is a working example, demonstrating integration between Go and ReactJS. You will find the relevant artefacts in the following folders:

* `./build/package/goreact` -- containing a series of docker based build scripts
* `./cmd/goreact` -- main package for app name `goreact`
* `./internal/goreact` -- basic packages for cli and server-side code
* `./scripts/goreact` -- a collection of BASH scripts to help you initiate build processes
* `./web/reactjs` -- a collection of reactjs projects

Please follow these steps:

* STEP 1: At the top this page, click the green coloured button named `Use this template` and follow the instructions from Github
* STEP 2: Modify [go.mod](./go.mod)
* STEP 3: Replicate the `goreact` examples and customise a version that suits your need

## Copyright

Copyright 2020 [go-web] Authors