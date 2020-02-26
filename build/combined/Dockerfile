# Copyright 2020 Paul Sitoh
# 
# Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# nodedev phase
# In this phase, we pull artefacts to build ReactJS webapp for
# development
FROM node:10.16.3 as nodedev

COPY ./web/reactjs/package.json /opt/package.json

RUN cd /opt && \
    npm install --save-dev mocha chai && \
    npm install --save-dev webpack webpack-cli webpack-dev-server webpack-merge && \
    npm install --save-dev html-webpack-plugin copy-webpack-plugin uglifyjs-webpack-plugin webpack-merge && \ 
    npm install --save-dev @babel/core babel-loader @babel/preset-env @babel/preset-react @babel/polyfill && \
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

# nodeprod phase
# Extract the ReactJS artefacts from the nodedev phase,
# build a production version of web artefacts and placed it 
# in folder 
# ./opt/web
#    bundle.js
#    images/
#    index.html
# 
FROM node:10.16.3 as nodeprod

COPY --from=nodedev ./opt/package-lock.json /opt/package-lock.json
COPY --from=nodedev ./opt/package.json /opt/package.json
COPY --from=nodedev ./opt/node_modules /opt/node_modules
COPY --from=nodedev ./opt/webpack /opt/webpack
COPY --from=nodedev ./opt/.babelrc /opt/.babelrc
COPY --from=nodedev ./opt/images /opt/images
COPY ./web/reactjs/src /opt/src

RUN cd /opt && npm audit fix && npm run build

# gobuild phase
# Utilising a go packaging tool github.com/GeertJohan/go.rice
# the web artefacts is packaged into a file name rice-box.go.
# Go builder then generates a version for linux platform.
FROM golang:1.13.3 as gobuild

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./build/combined/go-rice.sh ./build/go-rice.sh
COPY --from=nodeprod /opt/public ./web

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

# Set build flags to enable it to build artefacts that
# can be packed into scratch image.
# CGO_ENABLED=0 disable cross platform build - i.e. no need for libc
RUN go get github.com/GeertJohan/go.rice/rice && \
    ./build/go-rice.sh && \
    go mod download && \
    env CGO_ENABLED=0 env GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/package/go-react ./cmd/go-react

# Pack artefacts from gobuild phase into scratch for deployment
FROM scratch

COPY --from=gobuild /opt/build/package/go-react /go-react

CMD ["/go-react"]