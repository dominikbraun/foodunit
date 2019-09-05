#!/bin/bash
cd client
yarn build
cd ..
go run cmd/server/main.go