#!/bin/bash

# Copyright 2020 [go-web] Authors
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

env GOOS=linux GOARCH=amd64 go build -o ./build/package/linux/${APP_NAME} ./cmd/goreact
env GOOS=darwin GOARCH=amd64 go build -o ./build/package/macOS/${APP_NAME} ./cmd/goreact
env GOOS=windows GOARCH=amd64 go build -o ./build/package/windows/${APP_NAME}.exe ./cmd/goreact