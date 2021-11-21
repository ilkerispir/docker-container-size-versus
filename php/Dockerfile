
FROM php:8.0-apache

RUN docker-php-ext-install -j "$(nproc)" opcache

WORKDIR /var/www/html
COPY . ./

RUN sed -i 's/80/8080/g' /etc/apache2/sites-available/000-default.conf /etc/apache2/ports.conf


RUN mv "$PHP_INI_DIR/php.ini-development" "$PHP_INI_DIR/php.ini"