#!/bin/bash

echo "Starting FoodUnit server in development mode"

cd ./cmd/server || exit
go run main.go --addr :9292 &

cd ./../../ui || exit
yarn serve
