#!/bin/bash

echo "This script runs the migration using the standard configuration:"
echo "- Connecting to 'foodunit' database at localhost with root:root"
echo "- Don't drop existing tables"
echo ""
echo "If you want to change these values, start the migration manually."
echo "Intended for development only."
echo ""

echo "Starting FoodUnit migration"
echo ""

cd ./cmd/migration || exit
go run main.go