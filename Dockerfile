FROM golang:latest
RUN cd /go/src && mkdir -p github.com/koinkoin-io/ git clone https://github.com/koinkoin-io/koinkoin.backend && cd koinkoin.backend
WORKDIR /go/src/github.com/koinkoin-io/koinkoin.backend
ENTRYPOINT ["/usr/bin/make", "run"]
