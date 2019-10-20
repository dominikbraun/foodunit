# Copyright 2019 The FoodUnit Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# FoodUnit 3 API server image (Development Version)

FROM golang:1.13

RUN mkdir -p /foodunit
WORKDIR /foodunit

# PORT specifies the port the server listens on.
ENV PORT 9292
ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

# Use the mod files in order to download all required modules.
RUN go mod download
COPY . .

# `go get` will download CompileDaemon into /go/bin, therefore
# it will be availaible as a CLI command.
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

# Run CompileDaemon, enabling Hot Reloading. It will observe a mounted
# directory and rebuild/restart the app when a file changes.
ENTRYPOINT CompileDaemon \
    -build="go build -o ./.target/foodunit-server ./cmd/server/main.go" \
    -command="./.target/foodunit-server --addr :${PORT}" \
    -log-prefix=false

EXPOSE ${PORT}