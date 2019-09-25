#!/bin/bash

echo "Building FoodUnit binary"

mkdir -p ./.target
go build -o ./.target/foodunit-server cmd/server/main.go

# ToDo: Run FoodUnit as a `foodunit` user
# echo "Preparing server execution"

# getent passwd foodunit || useradd -m -d /home/foodunit foodunit
# chown -R foodunit:foodunit ./.target/*

echo "Starting FoodUnit server in production mode"

cd ./.target || exit
chmod +x foodunit-server
./foodunit-server --addr :9292
