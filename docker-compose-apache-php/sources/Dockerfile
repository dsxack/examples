FROM php:7.4-apache

RUN apt-get update && apt-get install -y --force-yes \
        default-libmysqlclient-dev  \
        libpng-dev \
        libmcrypt-dev \
        libjpeg-dev \
        libfreetype6-dev \
    && docker-php-ext-install -j$(nproc) mysqli pdo mcrypt pdo_mysql mbstring \
    && docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ \
    && docker-php-ext-install -j$(nproc) gd \
    && docker-php-ext-enable mysqli pdo pdo_mysql gd mcrypt mbstring

RUN yes | pecl install xdebug \
    && echo "zend_extension=$(find /usr/local/lib/php/extensions/ -name xdebug.so)" > /usr/local/etc/php/conf.d/xdebug.ini \
    && echo "xdebug.remote_enable=on" >> /usr/local/etc/php/conf.d/xdebug.ini \
    && echo "xdebug.remote_autostart=off" >> /usr/local/etc/php/conf.d/xdebug.ini

RUN yes | echo "mbstring.func_overload=2" >> /usr/local/etc/php/conf.d/mbstring.ini \
    && echo "mbstring.internal_encoding=utf-8" >> /usr/local/etc/php/conf.d/mbstring.ini \
    && echo "date.timezone=Europe/Moscow" >> /usr/local/etc/php/conf.d/date.ini \
    && echo "session.save_path=/tmp" >> /usr/local/etc/php/conf.d/session.ini

COPY etc/apache2/sites-enabled/000-default.conf /etc/apache2/sites-enabled/000-default.conf

RUN a2enmod rewrite

CMD ["apache2-foreground"]
