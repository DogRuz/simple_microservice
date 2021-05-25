FROM golang:1.16.2

RUN go version
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool


WORKDIR /app
COPY ./ /app

#RUN go get -v -t  .v