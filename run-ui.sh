#!/bin/bash

echo "This script starts the UI using the standard configuration:"
echo "- Listening on port 3000"
echo ""
echo "If you want to change these values, start the server manually."
echo "Intended for development only."
echo ""

echo "Starting FoodUnit server"
echo ""

cd ./ui || exit
yarn serve