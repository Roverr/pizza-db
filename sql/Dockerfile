FROM mysql:latest
MAINTAINER eCDN

ENV MYSQL_DATABASE=pizzeria
COPY structure.sql /docker-entrypoint-initdb.d/schema.sql
