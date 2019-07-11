#
# FoodUnit core application
#

# Perform a multistage build with Composer and PHP.
FROM composer:1.8 AS composer
FROM php:7.1.27-apache

COPY --from=composer /usr/bin/composer /usr/bin/composer

# Move project files into the document root.
COPY . /srv/app/public

# Explicitly override the default configuration.
COPY .docker/app/vhost.conf /etc/apache2/sites-available/000-default.conf

# Install all required APT packages. Clear cache afterwards.
RUN apt-get update && apt-get install -y \
        git \
        zlib1g-dev \
        unzip \
    && rm -rf /var/lib/apt/lists/*

# Install all required PHP extensions.
# Set ownership of the document root.
RUN docker-php-ext-install \
        mbstring \
        pdo \
        pdo_mysql \
        zip \
    && chown -R www-data:www-data /srv/app/public && a2enmod rewrite

# Switch into the API directory and install all packages.
WORKDIR /srv/app/public/api

# do not run composer as root
USER www-data
RUN COMPOSER_CACHE_DIR=/dev/null composer install
USER root