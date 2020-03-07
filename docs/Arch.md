# Architecture

When you create scaffold, you'll find a skeleton ReactJS and Go based web/rest backend. This does not mean you are confined to building projects based on these technologies. You can easily replace the ReactJS source with VueJS or any other framework. You can easily extends the Go REST backend.

## Project layout

To support architecture goal, the layout of your scaffold is based on the [standard Go project layout](https://github.com/golang-standards/project-layout).

Out-of-the-box, you will find the following artefacts, and these are 

### `build`

This folder contains artefacts, mainly Dockerfiles, responsible for creating native apps or container images.

### `cmd`

This folder contains the source codes or Go main package for user facing aspect of you app namely your exectuables. This folder should only containts codes to initialise, starting and stopping your application process.

### `web`

This folder contains source codes, in the context of your initial scaffold ReactJS, for building the web app. You should only keep your webapp source code and its build artefacts here and no where else. 

### `internal`

This folder contains all the Go dependencies that are used to support Go packages in `cmd`. The Go packages in this folder are not exportable.

## Modifying/extending scaffold

TO-DO
