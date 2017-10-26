FROM alpine:3.4
MAINTAINER Imre Racz <rover@ustream.tv>
RUN apk add --update bash openssl ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir /app
COPY pizza-db /app/
CMD ["/pizza-db"]
