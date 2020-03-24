FROM alpine:3.10.2
MAINTAINER 242617@gmail.com

COPY build/core /usr/local/core
COPY build/config.yaml /etc/core/config.yaml

CMD /usr/local/core --config /etc/core/config.yaml