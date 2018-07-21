FROM golang:latest

RUN mkdir -p /go/src/github.com/koinkoin-io/koinkoin.backend
COPY . /go/src/github.com/koinkoin-io/koinkoin.backend
WORKDIR /go/src/github.com/koinkoin-io/koinkoin.backend

ENTRYPOINT ["/usr/bin/make", "run"]
