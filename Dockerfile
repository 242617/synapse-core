FROM 242617/go-builder:1.0.0 AS builder

ARG PROJECT
ARG APPLICATION
ARG ENVIRONMENT
ARG VERSION

ENV PROJECT=${PROJECT}
ENV APPLICATION=${APPLICATION}
ENV ENVIRONMENT=${ENVIRONMENT}
ENV VERSION=${VERSION}

WORKDIR /root
COPY . .
RUN make proto
RUN make build

FROM alpine:3.10.2

WORKDIR /usr/local
COPY --from=builder /root/build/core .
WORKDIR /etc/core
COPY build/config.yaml .

CMD /usr/local/core --config /etc/core/config.yaml