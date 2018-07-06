FROM golang:latest
RUN cd /go/src && git clone https://github.com/koinkoin-io/koinkoin.backend && cd koinkoin.backend
CMD ["/usr/local/go/bin/go run main.go"]