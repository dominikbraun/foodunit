## FoodUnit 3

### Prerequisites

Running FoodUnit 3 natively requires:
* Go 1.12: https://golang.org
* Yarn 1.17: https://yarnpkg.com

Running FoodUnit 3 containerized requires:
* Docker 19.03: https://docker.com

### Start FoodUnit natively

FoodUnit basically consists of an API server, an UI server, a migration tool and a backing database. In order to run the application natively, a local MySQL/MariaDB database is required.

Copy `app.example.toml` and rename it to `app.toml`. Fill in appropriate values. Do the same thing with `ui/public/config.example.json`.

If the database tables do not exist yet, run:

```shell script
$ ./run-migration.sh
```

Then start the other components:

```shell script
$ ./run-server.sh
$ ./run-ui.sh
```

If you want to change the default values, start the component manually with the corresponding flags. These scripts are intended for development only.


### Start FoodUnit containerized

Just as mentioned above, copy and rename the configuration files. Then build and run the containers.

```shell script
$ docker-compose up
```

You may run the containers yourself just with `docker run`, using the environment variables and volumes specified in `docker-compose.yml`.

These containers are intended for development only.
