#!/bin/bash

echo "This script starts the server using the standard configuration:"
echo "- Listening on port 9292"
echo "- Connecting to 'foodunit' database at localhost with root:root"
echo ""
echo "If you want to change these values, start the server manually."
echo "Intended for development only."
echo ""

echo "Starting FoodUnit server"
echo ""

cd ./cmd/server || exit
go run main.go --addr :9292
