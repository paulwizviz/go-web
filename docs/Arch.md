# Architecture

Your scaffold has been architected to help you, as a developer, create a native macOS, Linux and Windows app or a Docker-based container.

## Folders

Your scaffold include folders layout based on recommendations described in [standard Go project layout](https://github.com/golang-standards/project-layout).

### `./build`

This folder contains the following:

* [./build/package](../build/package) folder containers build scripts in the form of docker and docker-compose files to produce native macOS, Linux and Windows app or a Docker-based image;
* [./build/package/go-rice.sh](../build/package/go-rice.sh) script to pre-generate Go source files embedding the ReactJS artefacts in byte code form. DO NOT REMOVE THIS FILE.

If you wish to extend the build process please to do it in this folder [./build/package](../build/package).

When you execute the build scripts you will find a generated folder call `./build/native`, where you find native versions of executables (macOS, Linux and Windows).

### `./cmd`

This folder contains basic Go codes to create cli commands, subcommands, arguments and flags to manage the lifecycle (start, stop, configure, etc) of your built artefacts.

Any subfolders here with `main.go` are treated as the entry point of your built artefacts. Out-of-the-box `./cmd/goreact`, which you should rename to fit your requirement.

### `./deployments`

This folder contains a [./deployments/dev](../deployments/dev) to support you development effort. Here you will find the following:

* [./deployments/dev/docker-compose.yaml](../deployments/dev/docker-compose.yaml) to create a local development environment (see content of development environment);
* [./deployments/dev/react.dockerfile](../deployments/dev/react.dockerfile) to help `docker-compose.yaml` create your ui container;
* [./deployments/dev/go-run.sh](../deployments/dev/go-run.sh) to help `docker-compose.yaml` create REST container.

Create more folders to support your operations to deploy your build artefacts to an infrastructure/platform of your choice (e.g. Kubernetes, etc).

### `./internal`

This folder contains skeleton Go source codes for the body your native or container apps. Please retain the folder `internal/server` and the Go files. It is needed to generate you web and REST servers.

### `./web`

This folder contains a skeleton RectJS codes found in the sub-folder `./web/reactjs`.

If you wish to work with other Javascript UI framework create an appropriate subfolder. For example, `./web/vue` for Vue framework, etc. You will need to make changes in the built scripts in `./build/package`.

### `./scripts`

This folder containers Bash scripts to trigger build processes and to execute various deployment scenarios. Out-of-the-box scaffold you will find the follow:

* [./scripts/container.sh](../scripts/container.sh) to trigger the build process for docker-images;
* [./scripts/dev.sh](../scripts/dev.sh) to trigger the build process for apps configure for development activities;
* [./scripts/native.sh](../scripts/native.sh) to trigger build processes to generate container apps.

### `./test`

This folder contains minimum resources (dockerfiles and bash scripts) specifically to support your testing effort.

## Development Environment

Your scaffold contains a development environment based on [docker-compose.yaml](../deployments/docker-compose.yaml) found in deployment folder. The `docker-compose` generates the following containers:

* An nginx-styled reverse proxy container based on this Docker image [implementation](https://github.com/binocarlos/noxy);
* A ReactJS container with code hot loading to enable you to modify your code and see changes immediately;
* A Go based REST container.
