# Architecture

Your scaffold has been architected to help you, as a developer, create a native macOS, Linux and Windows app or a Docker-based container.

## Scaffold folders

Your scaffold include folders layout based on recommendations described in [standard Go project layout](https://github.com/golang-standards/project-layout).

### `./build`

This folder contains the following:

* [./build/package/artefacts](../build/package/artefacts) folder containers Dockerfiles to produce native macOS, Linux and Windows app or a Docker image;

* [./build/package/go-rice.sh](../build/package/go-rice.sh) script to pre-generate Go source files embedding the ReactJS artefacts in byte code form. DO NOT REMOVE THIS FILE.

If you wish to extend the build process please to do it in this folder [./build/package](../build/package)

This will folder will also create a folder named `native`, which is git ignored, when you generate native apps. You will find the generated native apps for macOS, Linux and Windows here.

### `./cmd`

This folder contains basic Go codes to create cli commands, subcommands, arguments and flags to manage the lifecycle (start, stop, configure, etc) of your built artefacts.

Any subfolders here with `main.go` are treated as the entry point of your built artefacts. Out-of-the-box `./cmd/goreact`, which you should customised to fit your requirement, serve as the entry point to the native app wrapping Web and a REST service in one package. 

You will also find the artefact `./cmd/gorest` which will be use to generate a REST server to support you in the development phase. Keep this intact.

### `./deployments`

This folder contains basic `docker-compose` files to support:

* [./deployments/dev/docker-compose.yaml](./deployments/dev/docker-compose.yaml) development work.

* [./deployment/e2e/docker-compose.yaml](../deployment/e2e/docker-compose.yaml) container based end-to-end testing.

### `./internal`

This folder contains skeleton Go source codes for the body your native or container apps. Please retain the folder `internal/server` and the Go files. It is needed to generate you web and REST servers.

### `./web`

This folder contains a skeleton RectJS codes found in the sub-folder `./web/reactjs`.

If you wish to work with other Javascript UI framework create an appropriate subfolder. For example, `./web/vue` for Vue framework, etc. You will need to make changes in the built scripts in `./build/package`.

### `./scripts`

This folder containers Bash scripts to trigger build processes and to execute various deployment scenarios. Out-of-the-box scaffold you will find the follow:

* [./scripts/artefacts](../scripts/artefacts) to trigger the build process for production native apps;

* [./scripts/dev](../scripts/dev) to trigger the build process for apps configure for development activities;

* [./scripts/test](../scripts/test) to trigger test processes (unit and e2e);

* `./common.sh` - central location for your to customise your built artefacts.

### `./test`

This folder contains minimum resources (dockerfiles and bash scripts) specifically to support your testing. You will find [scripts](../test/unit) to support ReactJS and Go unit test. When you plan to extend your test frameworks (smoke tests, end-to-end, etc), place your support resources (mocks, etc) here.

## Development Environment

Your scaffold contains a development environment based on [docker-compose.yaml](../deployments/docker-compose.yaml) found in deployment folder. The `docker-compose` generates the following containers:

* An nginx-styled reverse proxy container based on this Docker image [implementation](https://github.com/binocarlos/noxy);

* A ReactJS container created from this [./build/dev/react/Dockerfile](../build/dev/react/Dockerfile) with code hot loading to enable you to modify your code and see changes immediately;

* A Go based REST container created from this [./build/dev/rest/Dockerfile](../build/dev/rest/Dockerfile).
