# Your App in Action

The following steps assumed that you have not customise your scaffold and using it as-is.

## Native apps

To see a native app in action, follow these steps:

1. Open an terminal and navigate to `./build/native/<platform of your choice>`;

2. Execute the command `goreact` (for macOS and Linux) or `goreact.exe` (for Windows);

**NOTE:** `goreact` is the name of the built artefact out-of-the-box.

3. Open your browser with the URL `localhost` or `localhost:<your-choice>`;

**NOTE:** By default the app runs on port 80. Type the command `goreact -port=<your-choice>` to run on a port of your choice.

## Container app

1. Execute the following command `./scripts/test/e2e.sh start`. This will start-up a container with port 80 open.

**NOTE:** If you wish to have the container expose a different port modify the docker-compose deployment found here [script](../deployments/e2e/docker-compose.yaml)

2. Assuming you have no problem with STEP 1 and operating on port 80, open a browser with url to `localhost`.

3. To stop you app, execute the following command `./scripts/test/e2e.sh stop`.

4. To clean up your workspace, `./scripts/test/e2e.sh clean`.

