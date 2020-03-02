# Combined Webapp (Javascript) and Go app template

This a template to help you scaffold a combined Webapp and Go project. The Webapp powers the UI interface and Go powers the RESTful API as well as providing hooks to start the combined app.

This is template is intended for projects involving two different sets of expertises namely, Javascript/NodeJS and Go.

**NOTE** The scaffold generated from this project whilst functional is not a full featured and definitely **NOT** for production or any mission critical use.

## Prerequisite

In order to build apps based on this template ensure you have the following items installed:

1. Install Docker and Docker compose

2. Install [Go toolkit version 1.11 onwards](https://blog.golang.org/)

3. (Optional) Install NodeJS (out-of-the-box you don't need this because this template uses Docker to build the node part)

## How to scaffold a project from this template?

Click on the template button and follow the instructions from Github.

## What do I get out of the box?

When you create a scaffold from this template you get the following:

1. A functional proxy and baseline web and API server for you to extend from.

2. An example ReactJS frontend that you can extend or draw lesson.

3. A environment to support web development.

4. Build scripts -- Docker based -- to create native (macOS, Linux and Windows) app or container based app.

5. The scaffold is based on Go modules (i.e. Go11+)

## How is my scaffold layout?

Your scaffold follows the [Go project layout](https://github.com/golang-standards/project-layout). In your scaffold you will find the follow:

### `build`

This folder contains artefacts, mainly Dockerfiles, to responsible for creating native apps or container images.

### `cmd`

This folder contains the source codes or Go main package for user facing aspect of you app namely your exectuables. This folder should only containts codes to initialise, starting and stopping your application process.

### `web`

This folder contains source codes, in the context of your initial scaffold ReactJS, for building the web app. You should only keep your webapp source code and its build artefacts here and no where else. 

### `internal`

This folder contains all the Go dependencies that are used to support Go packages in `cmd`. The Go packages in this folder are not exportable.
