# Combined ReactJS and Golang app

This project demonstrates a desktop executable browser based app built from a combination of ReactJS and Go codes. ReactJS powers the UI interface and Go powers the RESTful API. 


Please refer to [Dockerfile](./build/Dockerfile) where you will see the build process. Study the build process, which also applies to creating other platform (e.g. macOS, Windows, etc) based apps. Please refer to Go instruction on how to build for [multiple platforms](https://binx.io/blog/2018/11/25/go-cross-compilation/).

**IMPORTANT:** This is not a production ready nor a real working application that you can use out of the box to fulfil some real world use case. It is purely to demonstrate concept or use it as a template to create similar projects.

## Why ReactJS and Golang?

There is no particular reason. Having said that, an argument could be made on the grounds that this project is essentially a web stack; either language could have accomplished both goals on their own.

However, this particular use case is to demonstrate a scenario where a developer maybe more competent at writing backend stuff in Go and needing help from more competent UI developers (where there is more experience developers from the Javascript world). 

The project also serve to demonstrate that following the recommended [Go project layout](https://github.com/golang-standards/project-layout) allow for two different code-based to exists coherently in a single repository.

## Prerequisite

You are expected to have working knowledge of Go and ReactJS/Redux.

In order to run the project out-of-the-box, please ensure you have the following items installed:

1. Install Docker and Docker compose
2. Install [Go toolkit version 1.11 onwards](https://blog.golang.org/)
3. (Optional) Install NodeJS (out-of-the-box you don't need this because this project uses Docker to build the node part)

## Running the project out-of-the-box

1. Run `git clone http://github.com/paulwizviz/go-react` or alternatively in an appropriate folder of your choice and assuming you have already installed Go toolkit, run the command `go get github.com/paulwizviz/go-react`.

2. Navigate into the project folder, run the following command to build the project.
```
    ./scripts/build.sh package <version number e.g. 0.0>
```
You will find native apps for Linux, macOS and Windows in this location:
```
    ./build/package
```

3. To see the app in action, select the a version appropriate for your platform. Assuming you are on macOS, to execute your app, run the following command:
```
    ./build/package/macOS/go-react
```
You should see the following information in your console:
```
2020/02/02 17:19:50 Starting react 3000
2020/02/02 17:19:50 Starting rest server 0.0.0.0:9000
```

4. The app is based on the following default ports ReactJS (3000) and API (9000). To see that the app is running open your browser with the following url `localhost:3000` and, if no error, you will see a dashbaord like UI.

5. In the sidebar select the option `users` and you will see in the main dashboard a simple button labelled `FETCH USER`. Click on the button and if all goes well you will see:

```
{"data":"{users: [{\"name\":\"Albert\"}.{\"name\":\"Beatrice\"}]}","status":200,"statusText":"OK","headers":{"content-length":"48","content-type":"application/json"},"config":{"url":"http://localhost:9000/users","method":"get","headers":{"Accept":"application/json, text/plain, */*"},"transformRequest":[null],"transformResponse":[null],"timeout":0,"xsrfCookieName":"XSRF-TOKEN","xsrfHeaderName":"X-XSRF-TOKEN","maxContentLength":-1},"request":{}}
```

If you reached this far it means you have a fully working out-of-the-box app.

NOTE:

1. The Go RESTful API is implemented [here](./internal/rest/userhandler.go). This is a simple REST interface delivering a hardcoded set of users as shown in step 6 above.

2. The ReactJS-Redux implementation is found [here](./web/reactjs).

## What can you do with this project?

Most important takeaway from this project is to study the way the codes are organised and built. I suggest you study the Docker files in the `build` folder to see how things works. Once you understand the principles behind the project, experiment. 

If you have any queries, just raise an issue in the repo or feel free to PR.