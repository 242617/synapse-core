FROM alpine:3.10.2

WORKDIR /usr/local
COPY build/synapse .
WORKDIR /etc/synapse
COPY build/config.yaml .

CMD /usr/local/synapse --config /etc/synapse/config.yaml