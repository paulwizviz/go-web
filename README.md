# Combined Webapp (Javascript) and Go app template

This a template to help you scaffold a combined Webapp and Go project. The Webapp powers the UI interface and Go powers the RESTful API as well as providing hooks to start the combined app.

This is template is intended for projects involving two different sets of expertises Javascript and Go.

**NOTE** The scaffold generated from this project whilst functional is not a full featured and definitely **NOT** for production or any mission critical use.

## What to expected from this template

1. The template layout is closely follow the [Go project layout](https://github.com/golang-standards/project-layout). The source code for webapp is located under the folder `web`. The Go app codes are located in `cmd` (end-user facing entry-point), `internal` (private packages) and `pkg` (exportable packages).

2. The default webapp scaffold is based on ReactJS UI framework but more (e.g. VueJS) could be added in future. If you already have an existing Web frontend you can easily add to this project. Simply replace the out-of-the-box ReactJS framework under the folder `web`.

3. The combined app is packaged as a single container which should allow you to run on any platform via Docker. Future development may include option to build native apps (macOS, Linux and Windows).

## How to scaffold a project from this template?

Click on the template button and follow the instructions from Github.

## Prerequisite

In order to run any project, out-of-the-box, generated from this template, please ensure you have the following items installed:

1. Install Docker and Docker compose

2. Install [Go toolkit version 1.11 onwards](https://blog.golang.org/)

3. (Optional) Install NodeJS (out-of-the-box you don't need this because this template uses Docker to build the node part)
