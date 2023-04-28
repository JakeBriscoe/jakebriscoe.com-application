FROM alpine:latest

RUN apk update && \
    apk add postgresql-client openssl && \
    apk add --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing kubectl

WORKDIR /app
COPY startup.sh .

CMD ["/bin/sh", "/app/startup.sh"]