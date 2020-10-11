# Your App in Action

The following steps assumed that you have not customise your scaffold and using it as-is.

## Native apps

To see a native app in action, follow these steps:

1. Open an terminal and navigate to `./build/native/<platform of your choice>`;
2. Execute the command `goreact start ui` (for macOS and Linux) or `goreact.exe` (for Windows);
**NOTE:** 
* `goreact` is the name of the built artefact out-of-the-box.
* Replace `goreact` with `<your-choice>` to meet your requirement

3. Open your browser with the URL `localhost` or `localhost:<your-choice>`;
**NOTE:** 
* By default the app runs on port 80. 
* Type the command `goreact -port=<your-choice>` to run on a port of your choice.

## Container app

To verify that you have built your container accordingly, follow these steps:

1. Run the command `./scripts/container.sh run`;
2. Assuming you have no problem with STEP 1 and operating on port 80, open a browser with url to `localhost`;
3. To stop you app, execute the following command `./scripts/container.sh stop`.
4. To clean up your workspace, `./scripts/container.sh clean`.

## Automated test

Follow these steps to activate test scripts

1. Run the command `./scripts/test.sh unit`
2. Extend test accordingly.

## Copyright

Copyright 2020 [go-web] Authors