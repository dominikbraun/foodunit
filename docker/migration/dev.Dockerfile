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
# To build an image an run an instance manually, check out the docker-compose.yml file
# and use the corresponding volumes, environments variables, networks etc.
#
# This image can be built with `docker image build -t devmig -f docker/migration/dev.Dockerfile .`.
# To execute a migration, at least the database container has to be running. If you've started the
# database using the docker-compose.yml file, just run
# `docker container run --name foodunit_mig_1 --rm --network foodunit_dn -e DSN="root:root@(172.18.0.2:3306)/foodunit?parseTime=true" migration`.

# Start build stage
FROM golang:1.12-alpine

# Add git since some Go modules may use `git init`.
RUN apk add git --no-cache

# DSN defines the data source name in the form "user:pass@(host:port)/dbName?parseTime=true
ENV DSN "root:root@(localhost:3306)/foodunit?parseTime=true"
# DROP indicates if the table schema including all rows should be dropped ("true") or not ("false").
ENV DROP "false"
ENV GO111MODULE=on

RUN mkdir -p /migration
WORKDIR /migration

COPY go.mod .
COPY go.sum .

# # Use the mod files in order to download all required modules.
RUN go mod download
COPY . .

# Run the migration and drop the table schema if necessary.
CMD if [ "$DROP" = "true" ]; then \
        go run cmd/migration/main.go --dsn ${DSN}; \
    else \
        go run cmd/migration/main.go --dsn ${DSN} --drop-schema; \
    fi