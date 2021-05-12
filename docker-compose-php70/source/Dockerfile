FROM php:7.3.27-fpm

RUN apt-get update && apt-get install -y --force-yes \
        libmysqlclient-dev  \
    && docker-php-ext-install -j$(nproc) mysqli pdo pdo_mysql \
    && docker-php-ext-enable mysqli pdo pdo_mysql

CMD ["php-fpm"]
