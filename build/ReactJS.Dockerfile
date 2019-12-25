# Development
FROM node:10.16.3 as development

COPY ./web/reactjs/package.json /opt/package.json

RUN cd /opt && \
    npm install --save-dev mocha chai && \
    npm install --save-dev webpack webpack-cli webpack-dev-server webpack-merge && \
    npm install --save-dev html-webpack-plugin copy-webpack-plugin uglifyjs-webpack-plugin webpack-merge && \ 
    npm install --save-dev @babel/core babel-loader @babel/preset-env @babel/preset-react && \
    npm install --save-dev http-server && \
    npm install --save-dev error-overlay-webpack-plugin && \
    npm install --save-dev eslint-loader eslint eslint-plugin-react && \
    npm install --save-dev babel-eslint && \
    npm install --save react react-dom react-redux react-router-dom && \
    npm install --save redux redux-logger redux-thunk redux-promise-middleware && \
    npm install --save @material-ui/core @material-ui/icons && \
    npm install --save history && \
    npm install --save prop-types &&\
    npm install --save classnames && \
    npm install --save clsx && \
    npm install --save axios 

COPY ./web/reactjs/webpack /opt/webpack
COPY ./web/reactjs/.babelrc /opt/.babelrc
COPY ./web/reactjs/images /opt/images
COPY ./web/reactjs/.eslintrc.json /opt/.eslintrc.json

CMD ["npm","run","dev:run"]

# Production
FROM node:10.16.3 as production

COPY --from=development ./opt/package.json /opt/package.json
COPY --from=development ./opt/webpack /opt/webpack
COPY --from=development ./opt/.babelrc /opt/.babelrc
COPY --from=development ./opt/node_modules /opt/node_modules
COPY --from=development ./opt/images /opt/images
COPY ./web/reactjs/src /opt/src

RUN cd /opt && npm run build

# Release
FROM node:10.16.3

COPY --from=production /opt/public /opt/public
COPY ./web/reactjs/package.json /opt/package.json
RUN cd /opt && npm install --save-dev http-server

CMD ["npm", "start"]