FROM php:7.1.8-apache

COPY . /srv/app/public
COPY .docker/php/vhost.conf /etc/apache2/sites-available/000-default.conf

RUN chown -R www-data:www-data /srv/app/public && a2enmod rewrite
