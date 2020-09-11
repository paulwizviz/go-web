# Architecture

Your scaffold has been architected to help you, as a developer, create a native macOS, Linux and Windows app or a Docker-based container.

## Customising the scaffold

1. Modify the Go module name in `go.mod` to suite your requirement; 

2. Modify the folder name `./cmd/goreact` to one of your choice `./cmd/<youur-choice>`

3. Modify the content `./scripts/common.sh`

## Scaffold folders

Your scaffold has the following folders you should retain. 

These folders are based on recommendations described in [standard Go project layout](https://github.com/golang-standards/project-layout).

### `./build`

This folder contains the following:

* [./build/package/dev](../build/package/dev) folder includes Dockerfiles to build artefacts to support your development effort;

* [./build/package/artefacts](../build/package/artefacts) folder containers Dockerfiles to produce native macOS, Linux and Windows app or a Docker image;

* [./build/package/go-rice.sh](../build/package/go-rice.sh) script to pre-generate Go source files embedding the ReactJS artefacts in byte code form. DO NOT REMOVE THIS FILE.

If you wish to extend the build process please to do it in this folder [./build/package](../build/package)

This will folder will also create a folder named `native`, which is git ignored, when you generate native apps. You will find native apps for macOS, Linux and Windows.

### `./cmd`

This folder contains basic Go codes to manage the lifecycle (i.e. start and stop) of your build artefacts. Extend the codes to meet your commandline requirement requirements (e.g. adding command line to interact with your running code).

### `./deployments`

This folder contains basic `docker-compose` files to support:

* [./deployments/dev/docker-compose.yaml](./deployments/dev/docker-compose.yaml) development work.

* [./deployment/e2e/docker-compose.yaml](../deployment/e2e/docker-compose.yaml) container based end-to-end testing.

### `./internal`

This folder contains skeleton Go source codes for the body your native or container apps. Please retain the folder `internal/server` and the Go files.

### `./web`

This folder contains a skeleton RectJS codes found in the sub-folder `./web/reactjs`.

If you wish to work with other Javascript UI framework create an appropriate subfolder. For example, `./web/vue` for Vue framework, etc.

### `./scripts`

This folder containers Bash scripts to trigger build processes and to execute various deployment scenarios. In your scaffold you will find the follow:

* [./scripts/artefacts](../scripts/artefacts) to trigger the build process for production native apps;

* [./scripts/dev](../scripts/dev) to trigger the build process for apps configure for development activities;

* [./scripts/test](../scripts/test) to trigger test processes (unit and e2e);

* `./common.sh` - modify the to suit your requirement. 

## Development Environment

Your scaffold contains a development environment based on [docker-compose.yaml](../deployments/docker-compose.yaml) found in deployment folder. The `docker-compose` generates the following containers:

* An nginx-styled reverse proxy container based on this Docker image [implementation](https://github.com/binocarlos/noxy);

* A ReactJS container created from this [./build/dev/react/Dockerfile](../build/dev/react/Dockerfile) with code hot loading to enable you to modify your code and see changes immediately;

* A Go based RESTFul container created from this [./build/dev/rest/Dockerfile](../build/dev/rest/Dockerfile).
