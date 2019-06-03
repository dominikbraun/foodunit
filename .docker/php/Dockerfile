FROM php:7.1.8-apache

COPY . /srv/app/public

# overwrite default vhost to ensure it is the only one
COPY .docker/php/vhost.conf /etc/apache2/sites-available/000-default.conf

WORKDIR /srv/app/public

# install PHP extensions for MySQL
RUN docker-php-ext-install mbstring pdo pdo_mysql

# set ownership and enable rewriting
RUN chown -R www-data:www-data /srv/app/public && a2enmod rewrite
