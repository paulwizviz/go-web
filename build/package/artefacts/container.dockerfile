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

# node build phase
# In this phase, we pull artefacts to build ReactJS webapp for
# development
FROM node:13.10.1 as npminstall

WORKDIR /opt

COPY ./web/reactjs/dep.sh ./dep.sh
COPY ./web/reactjs/package.json ./package.json

RUN ./dep.sh

FROM node:13.10.1 as nodebuild

WORKDIR /opt

COPY --from=npminstall /opt/node_modules ./node_modules
COPY --from=npminstall /opt/package-lock.json ./package-lock.json
COPY --from=npminstall /opt/package.json /opt/package.json
COPY ./web/reactjs/webpack /opt/webpack
COPY ./web/reactjs/.babelrc /opt/.babelrc
COPY ./web/reactjs/images /opt/images
COPY ./web/reactjs/src /opt/src

RUN npm run build

# Go build phase
# Utilising a go packaging tool github.com/GeertJohan/go.rice
# the web artefacts is packaged into a file name rice-box.go.
# Go builder then generates a version for linux platform.
FROM golang:1.13.3 as gobuild

WORKDIR /opt

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./build/package/go-rice.sh ./build/go-rice.sh
COPY --from=nodebuild /opt/public ./web

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

# Replace app name {./cmd/goreact} here with name of your choice {./cmd/<your-choice>}
RUN go get github.com/GeertJohan/go.rice/rice && \
    ./build/go-rice.sh && \
    go mod download && \
    env CGO_ENABLED=0 env GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/package/container/goreact ./cmd/goreact

# Pack linux artefact into scratch container
FROM scratch

# Replace app name {goreact} here with name of your choice
COPY --from=gobuild /opt/build/package/container/goreact /goreact

# Replace app name {goreact} here with name of your choice
CMD ["/goreact"]