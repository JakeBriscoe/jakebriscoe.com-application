FROM alpine:latest

RUN apk update && \
    apk add postgresql-client kubectl openssl

WORKDIR /app
COPY startup.sh .

CMD ["/bin/bash", "/app/startup.sh"]