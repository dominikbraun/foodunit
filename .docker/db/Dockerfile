#
# FoodUnit database
#

FROM mariadb:10.3.15

# Mount the directory in order to provide a set of SQL files to execute.
ADD .docker/sql /docker-entrypoint-initdb.d

# Set required env variables.
ENV MYSQL_DATABASE foodunit2
ENV MYSQL_USER app
ENV MYSQL_PASSWORD app
ENV MYSQL_ROOT_PASSWORD foodunit2

# Expose the default port. The port has to be explicitly published on docker run
# if you want to access the container from your local machine.
EXPOSE 9906
