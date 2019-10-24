## FoodUnit 3

### Prerequisites

Running FoodUnit 3 natively requires:
* Go 1.12: https://golang.org
* Yarn 1.17: https://yarnpkg.com

Running FoodUnit 3 containerized requires:
* Docker 19.03: https://docker.com

### Starting the API server

To start the API server in development mode, use

```shell script
$ go run ./cmd/server/main.go --addr :9292
```

or use the provided script `run-dev.sh` which will also start the frontend server.

### Starting the frontend server

This step isn't necessary if you've used `run-dev.sh`. Otherwise:

```shell script
$ cd ./ui
$ yarn serve
```
