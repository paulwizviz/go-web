#!/bin/bash

# Dev dependencies
## Babel
npm install --save-dev @babel/core@7.11.4
npm install --save-dev @babel/plugin-proposal-class-properties@7.10.4
npm install --save-dev @babel/plugin-proposal-export-default-from@7.10.4
npm install --save-dev @babel/polyfill@7.10.4
npm install --save-dev @babel/preset-env@7.11.0
npm install --save-dev @babel/preset-react@7.10.4
npm install --save-dev babel-eslint@10.1.0
npm install --save-dev babel-loader@8.1.0

## ESLint
npm install --save-dev eslint@7.7.0
npm install --save-dev eslint-loader@4.0.2
npm install --save-dev eslint-plugin-react@7.20.6

## Test
npm install --save-dev mocha@8.1.2
npm install --save-dev chai@4.2.0

## Webpack 
npm install --save-dev webpack@4.44.1
npm install --save-dev webpack-cli@3.3.12
npm install --save-dev webpack-dev-server@3.11.0
npm install --save-dev webpack-merge@5.1.2
npm install --save-dev html-webpack-plugin@4.3.0 # Plugins
npm install --save-dev copy-webpack-plugin@6.0.3
npm install --save-dev error-overlay-webpack-plugin@0.4.1

## Others
npm install --save-dev http-server@0.12.3
npm install --save-dev core-js@3

# Depedencies
## Materials UI
npm install --save @material-ui/core@4.11.0
npm install --save @material-ui/icons@4.9.1

## React
npm install --save react@16.13.1
npm install --save react-dom@16.13.1
npm install --save react-redux@7.2.1
npm install --save react-router-dom@5.2.0
npm install --save redux@4.0.5 # redux
npm install --save redux-logger@3.0.6
npm install --save redux-promise-middleware@6.1.2
npm install --save redux-thunk@2.3.0

## Others
npm install --save axios@0.20.0
npm install --save classnames@2.2.6
npm install --save clsx@1.1.1
npm install --save history@5.0.0
npm install --save prop-types@15.7.2

npm audit fix 
npm audit
cat ./package.json