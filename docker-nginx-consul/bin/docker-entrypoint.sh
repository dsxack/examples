#!/usr/bin/env bash

nginx -g "daemon off;" \
	& consul-template -template "/etc/nginx/conf.d/default.conf.ctmpl:/etc/nginx/conf.d/default.conf:nginx -s reload"
