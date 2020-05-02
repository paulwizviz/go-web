# Architecture

This project has been architect to support you, as developers, create and package your web UI and RESTful interfaces into a single macOS, Linux and Windows app or a single Docker container.

## Project layout

The layout of your scaffold is based on the [standard Go project layout][1].

Out-of-the-box, you will find the following folders and these are as follows:

### `./build`

This folder includes artefacts to help you build development and production artefacts.

* [./build/package/dev](../build/package/dev) folder includes Dockerfiles to create a container of the sample ReactJS UI found in this project and a container of Go based REST server to support development activities.

* [./build/package/artefacts](../build/package/artefacts) folder containers Dockerfiles to produce native macOS, Linux and Windows app or a Docker image for a containerisable version of your app.

* [./build/package/go-rice.sh](../build/package/go-rice.sh) script to pre-generate a Go source file embedding the ReactJS artefacts in byte code form.

A folder named `native`, not version controlled, will be generated containing native apps for macOS, Linux and Windows platform.

### `./cmd`

This folder includes source codes to create your app entrypoint (i.e. main). 

These executables is based on the following Go frameworks [Cobra](https://github.com/spf13/cobra)

### `./deployments`

This folder contains two types of deployment scripts:

1. To support local development activities, you will find the following [docker-compose.yaml](../deployments/docker-compose.yaml). This script will generate the following containers locally:

    1.1. An nginx-styled reverse proxy known as noxy based on this [implementation](https://github.com/binocarlos/noxy)

    1.2. A react container built from this project see [./build/dev/react/Dockerfile](../build/dev/react/Dockerfile) with features for hot javascript loading.

    1.3. A Go based RESTFul container built from this project see [./build/dev/rest/Dockerfile](../build/dev/rest/Dockerfile)

### `./internal`

This folder container source codes to build an embedded webserver package and RESTFul endpoints. The codes are packaged as follows:

* `server` containing codes for a RESTful and web servers;

* `usermgmt` containing codes for user management handler.

The name of this folder has a special impact on project scoping. Please refer to this [doc](https://blog.learngoprogramming.com/special-packages-and-directories-in-go-1d6295690a6b) for explanation before you modify it.

### `./web`

This folder contains a demonstration RectJS sub project where you have all the necessary artefacts to create a web UI.

### `./scripts`

This folder containers operational scripts to enable you trigger build processes and to execute various deployment scenarios.

## Development environment

This template provides a mini development environment to help you develop the UI element. During development, especially of the UI elements, you may not need to connect to an actual backend service but simply to mock APIs. A separate dev environment build script is provided so you can create mocks of backend. It is also setup to enable hot loading of, out-of-the-box, ReactJS code.

For examples of how you can mock backends for dev enviornment, please refer to the [dev build scripts](../build/package/dev/rest.dockerfile) and this [Go code](../internal/usersmgmt/handlers/authuser/authhandler.go) to get a sense of how you can mock RESTFul interfaces.

## Modifying/extending scaffold

* Keep the layout intact and if you need to extends try to follow the recommendations describe [here][1]. Most importantly, retain the folder `internal` for your Go codes that you are creating as part of your application offerings. If you wish to make your project codes externally available for others to `go get` create the folder `pkg` and place your codes there. Keep anything related to your final build artefacts under these two folders `internal` and or `pkg` only. If you are planning to build another component of a microservices architecture, the recommendation is to create it as a separate Go module under a different github or equivalent repository.

* If you have specialise codes like GRPC protobufs or even non Go and non-web subprojects (Java, Python, etc), feel free to create one at the top level. Please use appropriate names (e.g. `java` for Java subprojects).

* The recommendation is to use Docker as the basis to build your project. If you stick to this recommendation, all you need to do is to modify/add dockerfiles found in the [build/package][2] folder.

* You will need to modify the module name in [go.mod](../go.mod) to one that suits your project. Please use appropriate naming conventions particularly if you plan to make your Go codes externally or `go get` visible.

* You will need to modify the folder under the top level folder `cmd` to one that is appropriate for your project. Remember to change modify the build script [build/package][2] accordingly. All you need to do is to is to get `go build` to reference the folder `./cmd` and sub folder with containing the main function. For example `go build -o <target location> ./cmd/<folder to your main file>`

* Replace the sample codes under the folder `internal` and `web` to one appropriate for your project. If you wish to build from the sample, keep and modify accordingly.

* Modify the copyright notices appropriate for your project.

[1]: https://github.com/golang-standards/project-layout
[2]: ../build/package