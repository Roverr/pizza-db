FROM alpine:3.4
MAINTAINER Imre Racz <rover@ustream.tv>
RUN apk add --update bash openssl ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir /app
COPY pizza-db /app/
COPY public /public
CMD ["/app/pizza-db"]
