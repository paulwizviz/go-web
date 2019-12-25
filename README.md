# Combined ReactJS and Golang app

This project demonstrates a desktop executable browser based app built from a combination of ReactJS and Go codes. ReactJS powers the UI interface and Go powers the RESTful API. In this particular case, Docker is also used as a runtime environment. However, it is possible to build and run the app on your desktop natively. But I leave it to you to experiment.

Why ReactJS and Golang?

There is no particular reason. This project is simply to demonstrate how you could combined the two into a single but yet coherent project repo. Having said that, an argument could be made on the grounds that this project is essentially a web stack; either language could have accomplished both goals on their own.

However, in my particular experience, where I find myself more competent at writing backend stuff in Go and needing help from more competent UI developers (where there is more experience developers from the Javascript world), this approach might seemed like a good compromise for developing cross platform Desktop app. In this particular case, I could do the backend whilst working collaboratively with a Javascript based developers who will focus on the frontend.

NOTE: This is not a production ready nor a real working application that you can use out of the box to fulfil some real world use case. It is purely to demonstrate concept or use it as a template to create similar projects

## Prerequisite

You are expected to have Go coding knowledge, ideally, ReactJS/Redux and Docker.

Please also acquaint yourself with the recommended [Go project layout](https://github.com/golang-standards/project-layout) which this project is based on.

In order to run the project out-of-the-box, please ensure you have the following items installed:

1. Install Docker and Docker compose
2. Install [Go toolkit version 1.11 onwards](https://blog.golang.org/)
3. (Optional) Install NodeJS (out-of-the-box you don't need this because this project uses Docker to build the node part)

## Running the project out-of-the-box

1. Run `git clone http://github.com/paulwizviz/go-react` or alternatively in an appropriate folder of your choice and assuming you have already installed Go toolkit, run the command `go get github.com/paulwizviz/go-react`.

2. Navigate into the project folder, run the following command to build the project.
```
    ./scripts/devOps.sh build <give it an image tag>
```

3. In the same project location, run the following command to start the app so it runs locally.
```
    ./scripts/devOps.sh start <image you use to build the project>
```
This runs the ReactJS app in development mode.

4. The app is based on the following default ports ReactJS (30303) and API (9000). To see that the app is running properly, run the following command:
```
    docker ps -a
```
You should see the following something similar to this:
```
CONTAINER ID        IMAGE                         COMMAND                  CREATED             STATUS              PORTS                    NAMES
395b4de26ec8        paulwizviz/react-ui-dev:0.0   "docker-entrypoint.s…"   3 hours ago         Up 8 seconds        0.0.0.0:3030->3030/tcp   ui-dev
fd4927ff7e13        paulwizviz/go-api-dev:0.0     "bash -c /usr/local/…"   3 hours ago         Up 9 seconds        0.0.0.0:9000->9000/tcp   api-dev
```

5. Open your browser with the following url `localhost:3030` and, if no error, you will see a simple button labelled `FETCH USER`. Click on the button and if all goes well you will see:

```
{"data":"{users: [{\"name\":\"Albert\"}.{\"name\":\"Beatrice\"}]}","status":200,"statusText":"OK","headers":{"content-length":"48","content-type":"application/json"},"config":{"url":"http://localhost:9000/users","method":"get","headers":{"Accept":"application/json, text/plain, */*"},"transformRequest":[null],"transformResponse":[null],"timeout":0,"xsrfCookieName":"XSRF-TOKEN","xsrfHeaderName":"X-XSRF-TOKEN","maxContentLength":-1},"request":{}}
```

If you reached this far it means you have a fully working out-of-the-box app.

NOTE: This was app was built and tested in macOS (Catalina). You may need to adjust (especially the script `devOps.sh`) to accomodate for Linux or Windows (please build your own powershell or whatever appropriate shell).

## What can you do with this project?

Most important takeaway from this project is to study the way the codes are organised and built. I suggest you study the Docker files in the `build` folder to see how things works. Once you understand the principles behind the project, experiment. 

If you have any queries, just raise an issue in the repo or feel free to PR.