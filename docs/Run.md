# Your App in Action

You would have built your app in native or container form.

## Native apps

To see a native app in action, follow these steps:

1. Open an terminal and navigate to `./build/native/<platform of your choice>`

2. You have the following options:

    2.1. To run the app with a UI frontend, execute the command `goreact` (for macOS and Linux) or `goreact.exe` (for Windows).

    2.2. To run the app with no UI frontend (i.e. headless), execute the command `goreact start noui` (for macOS and Linux) or `goreact.exec start noui`

**NOTE:**

* `goreact` is the name of the built artefact out-of-the-box.

* By default the app will run require port 80 to be available. If you need the app to use other ports type the command `goreact -port=<your-choice>`.

3. Assuming you have no problem with STEP 2.1 and operating on port 80, open a browser with url to `localhost`.

## Container app

1. Execute the following command `./scripts/test/e2e.sh start`. This will start-up a container with port 80 open.

NOTE: If you wish to have the container expose a different port modify the docker-compose deployment found here [script](../deployments/e2e/docker-compose.yaml)

2. Assuming you have no problem with STEP 1 and operating on port 80, open a browser with url to `localhost`.

3. To stop you app, execute the following command `./scripts/test/e2e.sh stop`.

4. To clean up your workspace, `./scripts/test/e2e.sh clean`.

