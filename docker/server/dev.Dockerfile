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
# Build command: docker build -t srvdevimg -f docker/server/dev.Dockerfile .
# Run command: docker run --name srvdevctr --rm -p 9595:9292 -e PORT=9292 -e DSN="root:root@(172.18.0.2:3306)/foodunit?parseTime=true" -v ${pwd}:/foodunit --network funet srvdevimg

FROM golang:1.13

# PORT specifies the port the server listens on.
ENV PORT 9292
# DSN defines the data source name in the form "user:pass@(host:port)/dbName?parseTime=true
ENV DSN "root:root@(localhost:3306)/foodunit?parseTime=true"
ENV GO111MODULE=on

RUN mkdir -p /foodunit
WORKDIR /foodunit

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
    -build="go build cmd/server/main.go" \
    -command="./main --dsn ${DSN} --addr :${PORT}" \
    -directory=. \
    -exclude-dir=.git \
    -log-prefix=true

EXPOSE ${PORT}