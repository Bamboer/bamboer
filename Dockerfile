FROM ubuntu:latest

RUN apt-get update  && \
    apt-get -y install nginx

EXPOSE 80
ENTRYPOINT nginx -g "daemon off;"

