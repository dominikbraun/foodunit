#!/bin/bash
cd ui
yarn build
cd ..
go run cmd/fileserver/main.go