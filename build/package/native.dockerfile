# Copyright [go-web] Contributors
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

ARG NODE_VER
ARG GO_VER

FROM node:${NODE_VER} as npminstall

WORKDIR /opt

ARG WEB_FRAMEWORK

COPY ./web/${WEB_FRAMEWORK}/dep.sh ./dep.sh
COPY ./web/${WEB_FRAMEWORK}/package.json ./package.json

RUN ./dep.sh

FROM node:${NODE_VER} as nodebuild

WORKDIR /opt

ARG WEB_FRAMEWORK

COPY --from=npminstall /opt/node_modules ./node_modules
COPY --from=npminstall /opt/package-lock.json ./package-lock.json
COPY --from=npminstall /opt/package.json /opt/package.json
COPY ./web/${WEB_FRAMEWORK}/webpack /opt/webpack
COPY ./web/${WEB_FRAMEWORK}/.babelrc /opt/.babelrc
COPY ./web/${WEB_FRAMEWORK}/images /opt/images
COPY ./web/${WEB_FRAMEWORK}/src /opt/src

RUN npm run build

# Go build phase
# Utilising a go packaging tool github.com/GeertJohan/go.rice
# the web artefacts is packaged into a file name rice-box.go.
# Go builder then generates a version for linux platform.
FROM golang:${GO_VER}

ARG WEB_FRAMEWORK

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY --from=nodebuild /opt/public ./internal/server/${WEB_FRAMEWORK}

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

ARG APP_NAME

