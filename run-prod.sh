#!/bin/bash

echo "Building FoodUnit binary"

mkdir -p ./.target
go build -o ./.target/foodunit cmd/server/main.go

echo "Starting FoodUnit server in production mode"

cd ./.target || exit
./foodunit
