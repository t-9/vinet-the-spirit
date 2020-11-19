FROM golang:1.15.5-alpine3.12

ADD . /go/src/vinet

WORKDIR /go/src/vinet

RUN cd /go/src/vinet && \
    go build -o vinet