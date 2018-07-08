FROM golang:latest

RUN go get -u github.com/koinkoin-io/koinkoin.backend

WORKDIR /go/src/github.com/koinkoin-io/koinkoin.backend
ENTRYPOINT ["/usr/bin/make", "run"]
