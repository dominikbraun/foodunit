# Prerequisites
- yarn (plain npm may also work, not tested)

# How to run
The go server can either:
- serve the production files from the client
- just proxy another webserver serving the client files

The last option is needed if you run the client in debug mode via yarn. The proxy prevents a cross-origin error your webbrowser may throw.

## Production mode
You can use the provided script:
```
./run_prod.sh
```

What it does:

It "builds" the client:

```
yarn build
```

and then runs the go server which also serves the client:
```
go run cmd/server/main.go
```

This will serve the client directly. (default port is 9292)


## Development mode

You can use the provided scripts:
```
./vue_dev.sh
```
and
```
./run_dev.sh
```

The first one starts vue in dev-mode which serves the app to port 8080.
The second script starts the go server with the --client-url param.
This proxies the 8080 port to the go-server-port (e.g. 9292).

So you can should the client via http://localhost:9292

# VS Code run scripts
To debug Foodunit with vs code you can use the following run configurations:

```
{
            "type": "chrome",
            "request": "launch",
            "name": "vuejs: chrome",
            "url": "http://localhost:9292",
            "webRoot": "${workspaceFolder}/client/src",
            "breakOnLoad": true,
            "sourceMapPathOverrides": {
                "webpack:///./src/*": "${webRoot}/*",
                "webpack:///src/*": "${webRoot}/*",
                "webpack:///*": "*",
                "webpack:///./~/*": "${webRoot}/node_modules/*"
            }
        },
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "cwd": "${workspaceFolder}",
            "program": "${workspaceFolder}/cmd/server/main.go",
            "env": {},
            "args": [
                "--client-url=http://127.0.0.1:8080"
            ]
        }
```

Then you can
1. run ./vue_dev.sh via commandline
2. execute the debug-config "Launch"
3. execute the debug-config "vuejs: chrome"

This will enable you to debug both: vue-client and go-server

The first one can also be configured as task (tasks.json)

```
{
    // See https://go.microsoft.com/fwlink/?LinkId=733558
    // for the documentation about the tasks.json format
    "version": "2.0.0",
    "tasks": [
        {
            "label": "yarn-serve",
            "type": "shell",
            "command": "cd client && yarn serve",
            "isBackground": false,
        }
    ]
}
```