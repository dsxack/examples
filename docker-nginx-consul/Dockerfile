FROM nginx:latest
LABEL maintainer="Smotrov Dmitriy <dsxack@gmail.com>"

ENV CONSUL_TEMPLATE_VERSION 0.19.4

RUN apt-get update && \
	apt-get install -y wget

RUN cd /tmp && \
	wget "https://releases.hashicorp.com/consul-template/${CONSUL_TEMPLATE_VERSION}/consul-template_${CONSUL_TEMPLATE_VERSION}_SHA256SUMS" && \
	wget "https://releases.hashicorp.com/consul-template/${CONSUL_TEMPLATE_VERSION}/consul-template_${CONSUL_TEMPLATE_VERSION}_linux_amd64.tgz" && \
    sha256sum -c consul-template_${CONSUL_TEMPLATE_VERSION}_SHA256SUMS 2>&1 | grep OK && \
    tar -xzvf consul-template_${CONSUL_TEMPLATE_VERSION}_linux_amd64.tgz && \
    mv consul-template /bin/consul-template && \
    rm -rf /tmp

EXPOSE 80
EXPOSE 443

COPY default.conf.ctmpl /etc/nginx/conf.d/
COPY ./bin/docker-entrypoint.sh /bin/

ENTRYPOINT ["/bin/docker-entrypoint.sh"]
