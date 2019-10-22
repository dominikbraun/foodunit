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

# FoodUnit 3 migration image (Development Version)

# Start build stage
FROM golang:1.12-alpine

# Add git in order to clone the repository.
RUN apk add git --no-cache
RUN git clone https://github.com/dominikbraun/foodunit foodunit
WORKDIR foodunit

# Build the application without the symbol table and debug information.
RUN go build -v -ldflags="-s -w" -o ./.target/foodunit-migration ./cmd/migration/main.go

# Start final stage
FROM alpine:3.10

# Create the main application directory.
RUN mkdir -p /foodunit
WORKDIR /foodunit

# Copy the compiled binary from the build stage into the application directory
# as well as the application configuration file.
COPY --from=0 /go/foodunit/.target/foodunit-migration .
ADD app.toml .

# DSN defines the data source name in the form "user:pass@(host:port)/dbName?parseTime=true
ENV DSN "root:root@(localhost:3306)/foodunit?parseTime=true"
# DROP indicates if the table schema including all rows should be dropped ("true") or not ("false").
ENV DROP "false"

# Run the FoodUnit migration and drop the table schema if necessary.
CMD if [ "$DROP" = "true" ]; then \
        ./foodunit-migration --dsn ${DSN}; \
    else \
        ./foodunit-migration --dsn {$DSN} --drop-schema; \
    fi