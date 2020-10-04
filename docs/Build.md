# Build process

You scaffold contains scripts to build native and/or container-based apps. The scripts to trigger build process are located in the [scripts folder](../scripts).

## What happens when you build your project?

This happens:

1. Your web app, out-of-box ReactJS, is reduced to a combination of html, Javascript and other resources, and are placed in a folder named `public`.
2. A built script `./build/package/go-rice.sh` compress html, Javascript and resources, into Go source code containing byte codes and place it in the folder `internal/server`.
3. Your Go code is complied into native executables (macOS, Linux and Windows) embedding web and REST server, and html related byte codes.
4. If you opt to do so, your Linux executable is package into a docker container. 

**NOTE:** 

These steps occurs in Docker based built scripts (Please refer to `./build/package`).

## Build for development

The scaffold provides a mini development environment. Follow these steps to build, run, stop and clean your development environment.

1. Open terminal and navigate to the root of your scaffold
2. Run the command `./scripts/dev.sh run` to build and run container images for react and rest
```
CONTAINER ID        IMAGE                     COMMAND                  CREATED             STATUS              PORTS                    NAMES
80d4b66a92e6        paulwizviz/go-rest:dev    "/usr/local/bin/go-r…"   21 seconds ago      Up 19 seconds       0.0.0.0:9000->9000/tcp   rest
6536a2b8c41c        paulwizviz/go-react:dev   "docker-entrypoint.s…"   21 seconds ago      Up 19 seconds       0.0.0.0:3030->3030/tcp   react
c6c3bfb538ce        binocarlos/noxy           "bash /run.sh"           21 seconds ago      Up 19 seconds       0.0.0.0:80->80/tcp       router
```
3. Run the command `./scripts/dev.sh stop` to stop your running development environment.
4. Run the command `./scripts/dev.sh clean` to remove Docker images.

## Build for production

1. Open terminal and navigate to the root of your scaffold.

### Native macOS, Linux and Windows

2. Run the command `./scripts/native.sh build`.
3. You will find your built arfects in the folder `./build/native/<native platform>`.

### Docker container based app

2. Run the command `./scripts/container.sh build`. 
3. Run the command `docker images` and, for out-of-the-box unmodified sample, you will see the following Docker images listed in your local repository:
```
paulwizviz/go-react-container   current                 273d6b7a1b46        5 minutes ago       9.07MB
```

### Clean project

4. Run the command `./scripts/native.sh clean` to remove any built native artefacts or `./scripts/container.sh clean` to remove build docker images.
