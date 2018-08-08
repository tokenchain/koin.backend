FROM golang:latest

RUN mkdir -p /go/src/github.com/koin-bet/koin.backend
COPY . /go/src/github.com/koin-bet/koin.backend
WORKDIR /go/src/github.com/koin-io/koin.backend

ENTRYPOINT ["/usr/bin/make", "run"]
