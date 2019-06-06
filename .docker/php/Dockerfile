# Docker multistage build for the FoodUnit core application.

FROM composer:1.8 AS composer

FROM php:7.1.8-apache

COPY --from=composer /usr/bin/composer /usr/bin/composer

COPY . /srv/app/public

# overwrite default vhost to ensure it is the only one.
COPY .docker/php/vhost.conf /etc/apache2/sites-available/000-default.conf

WORKDIR /srv/app/public

# update the sources list for jessie - see Stack Overflow #55386246
RUN echo "deb [check-valid-until=no] http://archive.debian.org/debian jessie main" > /etc/apt/sources.list.d/jessie.list
RUN sed -i '/deb http:\/\/deb.debian.org\/debian jessie-updates main/d' /etc/apt/sources.list
RUN apt-get -o Acquire::Check-Valid-Until=false update

RUN apt-get update && apt-get install -y git zlib1g-dev
# install required PHP extensions
RUN docker-php-ext-install mbstring pdo pdo_mysql zip

# set ownership and enable rewriting
RUN chown -R www-data:www-data /srv/app/public && a2enmod rewrite

WORKDIR /srv/app/public/api

RUN composer install
