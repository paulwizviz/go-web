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

FROM node:13.10.1 as npminstaller

WORKDIR /opt

COPY ./web/reactjs/package.json /opt/package.json
COPY ./web/reactjs/dep.sh /opt/dep.sh

RUN /opt/dep.sh

FROM node:13.10.1

WORKDIR /opt

COPY --from=npminstaller /opt/package-lock.json /opt/package-lock.json
COPY --from=npminstaller /opt/package.json /opt/package.json
COPY --from=npminstaller /opt/node_modules /opt/node_modules
COPY ./web/reactjs/webpack /opt/webpack
COPY ./web/reactjs/.babelrc /opt/.babelrc
COPY ./web/reactjs/images /opt/images
COPY ./web/reactjs/.eslintrc.json /opt/.eslintrc.json

RUN npm install

CMD ["npm","run","dev:run"]