FROM golang:latest
RUN cd /go/src \
    mkdir -p github.com/koinkoin-io/ \
    cd github.com/koinkoin-io/ \
    git clone http://github.com/koinkoin-io/koinkoin.backend
WORKDIR /go/src/github.com/koinkoin-io/koinkoin.backend
RUN make install
ENTRYPOINT ["/usr/bin/make", "run"]
