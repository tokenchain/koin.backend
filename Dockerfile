FROM golang:latest
RUN cd /go/src && git clone https://github.com/koinkoin-io/koinkoin.backend && cd koinkoin.backend
ENTRYPOINT ["/usr/bin/make", "run"]
