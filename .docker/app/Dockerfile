#
# FoodUnit core application
#

# Perform a multistage build with Composer and PHP.
FROM composer:1.8 AS composer
FROM php:7.1.8-apache

COPY --from=composer /usr/bin/composer /usr/bin/composer

# Move project files into the document root.
COPY . /srv/app/public

# Explicitly override the default configuration.
COPY .docker/app/vhost.conf /etc/apache2/sites-available/000-default.conf

WORKDIR /srv/app/public

# The PHP image comes with a Debian version that contains an invalid Jessie Repository
# in its sources.list. This removes the invalid entry (see Stack Overflow #55386246).
RUN echo "deb [check-valid-until=no] http://archive.debian.org/debian jessie main" > /etc/apt/sources.list.d/jessie.list
RUN sed -i '/deb http:\/\/deb.debian.org\/debian jessie-updates main/d' /etc/apt/sources.list
RUN apt-get -o Acquire::Check-Valid-Until=false update

# Install all required APT packages.
RUN apt-get update && apt-get install -y git zlib1g-dev

# Install all required PHP extensions.
RUN docker-php-ext-install mbstring pdo pdo_mysql zip

# Set ownership of the document root.
RUN chown -R www-data:www-data /srv/app/public && a2enmod rewrite

# Switch into the API directory and install all packages.
WORKDIR /srv/app/public/api
RUN composer install
