# Build process

You scaffold contains scripts to build native and/or container-based apps. The scripts to trigger build process are located in the [scripts folder](../scripts).

## What happens when you build your project?

When you build your project, this happens:

1. Your web app, ReactJS in out-of-box scaffold, is reduced to a combination of html, Javascript and other resources, placed in a folder named public.

2. A built script `./build/package/go-rice.sh` converts html, Javascript and resources into source code containing byte codes.

3. Your Go code is complied into native executables (macOS, Linux and Windows) embedding web and REST server, and web artefacts.

4. If you opt to do so, your Linux executable is package into a docker container. 

**NOTE:** 

* These steps occurs in Docker based built scripts (Please refer to `./build/package`).

* Out-of-the-box, your built artefact is named `goreact` which is generated from `./cmd/goreact`.

## Build for development

The scaffold provides a mini development environment to support your development effort. Follow these steps to build, run, stop and clean your environment.

1. Open terminal and navigate to the root of your scaffold

2. Run the command `./scripts/dev/ops.sh build` to build container images for react and rest

3. Run the command `./scripts/dev/ops.sh run` to start a development environment comprising of three docker containers: router, frontend and rest server (see eaxmple below)
```
CONTAINER ID        IMAGE                     COMMAND                  CREATED             STATUS              PORTS                    NAMES
80d4b66a92e6        paulwizviz/go-rest:dev    "/usr/local/bin/go-r…"   21 seconds ago      Up 19 seconds       0.0.0.0:9000->9000/tcp   rest
6536a2b8c41c        paulwizviz/go-react:dev   "docker-entrypoint.s…"   21 seconds ago      Up 19 seconds       0.0.0.0:3030->3030/tcp   react
c6c3bfb538ce        binocarlos/noxy           "bash /run.sh"           21 seconds ago      Up 19 seconds       0.0.0.0:80->80/tcp       router
```

4. Run the command `./scripts/dev/ops.sh stop` to stop you development environment.

5. Run the command `./scripts/dev/ops.sh clean` to remove Docker images.

## Build for production

1. Open terminal and navigate to the root of your scaffold

### Native macOS, Linux and Windows

2. Run the command `./scripts/artefacts/build.sh native`

3. You will find your built arfects in the folder `./build/native/<native platform>`

### Docker container based app

2. If do not plan to customise, please jump to STEP 3. If you plan to customise edit the script `./scripts/common.sh` and modify the variable `$APP_IMAGE_NAME` and `$APP_IMAGE_TAG` to something that suits your project requirements.

3. Run the command `./scripts/artefacts/build.sh container`.

4. Run the command `docker images` and, for out-of-the-box unmodified sample, you will see the following Docker images listed in your local repository:
```
paulwizviz/go-react-container   current                 273d6b7a1b46        5 minutes ago       9.07MB
```

### Clean project

5. Run the command `./scripts/artefacts/build.sh clean`
