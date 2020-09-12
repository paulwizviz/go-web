![build](https://github.com/paulwizviz/go-web/workflows/build/badge.svg)

# Overview

Use this template to scaffold a project for creating macOS, Linux and Windows native app or Docker base app incorporating a Web-based frontend (ReactJS, Vue, etc) with Go bases backend services (Web server, REST server and etc).

## Prerequisite

You **must** install Docker and Docker compose as these toolkits are needed to build native apps and/or container-based app from the scaffold plus your extensions.

You will also need to install these artefacts:

* [Go toolkit version 1.11 onwards](https://blog.golang.org/)

* Node

These are used to help you when you are editing codes in your scaffold. However, these are not to build your native and/or container apps.

## <a name="out-of-the-box">What do I get out-of-the-box when I create a scaffold</a>?

You will get the following:

* A layout based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout);

* A ReactJS-based skeleton sub-project;

* A skeleton Go source code to generate webserver and RESTful server;

* Docker based build scripts to help you create native (macOS, Linux and Windows) app or Docker image app.

* A locally deployable development environment to help you see the interactions between your web source code and Go RESTFul API;

* Docker-based testing framework.

## How do I use this template?

If you work with the [out-of-the-box](#out-of-the-box) artefacts, please follow these steps:

* STEP 1: At the top this page, click the green coloured button named `Use this template` and follow the instructions from Github;

* STEP 2: Modify [go.mod](./go.mod);

* STEP 3: Rename this folder [./cmd/goreat] to one of your choice (e.g. `./cmd/<your-choice>`) and modify the content of this file [./scripts/common.sh](./scripts/common.sh);

* STEP 4: Extends the web code in [./web/reactjs](./web/reactjs);

* STEP 5: Extend and add Go codes in `./cmd/<your-choice>` and [./internal](./internal);

* STEP 6: Build your project, please refer to [build doc](./docs/Build.md);

* STEP 7: To see you built artefacts (native or container) apps, please refer to [running the app](./docs/Run.md) doc.

## How do I extend/modify my scaffold?

 You can modify the [out-of-the-box](out-of-the-box) artefacts (folders, bash scripts, Dockerfiles, docker-compose file, etc). Please consult [architecture](./docs/Arch.md) docs for further information.

 
