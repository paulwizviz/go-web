![build](https://github.com/paulwizviz/go-web/workflows/build/badge.svg)

# Overview

This is a template to help developers scaffold a project for creating macOS, Linux and Windows native app or Docker base app incorporating a Web-based frontend (ReactJS, Vue, etc) with Go bases backend services (Web server, REST server and etc).

What problem is this template meant to solve?

Often Go developers are faced with a need to provide a user interface to support backend services. For Go developers that do not want spend time writing frontend codes, they would like to be able to just pick-up pre-developed user interfaces and just integrate as part of their Go project. Web based user interfaces are prevalent that could be reused. Hence, the purpose of this template is to help Go developers quickly provide a web based user interace for their deskop and container based project.

## Prerequisite

You **must** install Docker and Docker compose as these toolkits are needed to build native apps and/or container-based app from the scaffold plus your extensions.

You will also need to install these artefacts:

* [Go toolkit version 1.11 onwards](https://blog.golang.org/)

* Node (version depends on the web framework you are using)

These are used to help you when you are editing codes in your scaffold. However, these are not to build your native and/or container apps.

## <a name="out-of-the-box">What do I get out-of-the-box when I create a scaffold</a>?

When you derive a project directly from this template, you will get the following:

* A layout based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout);

* A ReactJS-based skeleton sub-project;

* A skeleton Go source code to generate webserver and REST server;

* Docker based build scripts to help you create native (macOS, Linux and Windows) app or Docker image app.

* A locally deployable development environment to help you see the interactions between your web source code and Go REST API (including hot code reloading whilst you develop);

* Docker-based testing framework.

## How do I use this template?

Please follow these steps:

* STEP 1: At the top this page, click the green coloured button named `Use this template` and follow the instructions from Github;
* STEP 2: Modify [go.mod](./go.mod);
* STEP 3: Rename this folder [./cmd/goreat](./cmd/goreact) to one of your choice (`./cmd/<your-choice>`);
* STEP 4: Modify the following scripts [./build/package/builder.yaml](./build/package/builder.yaml) and [./deployments/dev/docker-compose.yaml](./deployments/dev/docker-compose.yaml);
* STEP 5: Extends the web code in [./web/reactjs](./web/reactjs);
* STEP 6: Extend and add Go codes in `./cmd/<your-choice>` and [./internal](./internal);
* STEP 7: Build your project, please refer to [build doc](./docs/Build.md);
* STEP 8: To see you built artefacts (native or container) apps in action, please refer to [running the app](./docs/Run.md) doc.

## Advance customisation?

Please consult these docs for further information:

* [Architecture doc](./docs/Arch.md);
* [Build doc](./docs/Build.md);
* [Run your app](./docs/Run.md)