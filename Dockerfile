FROM golang:latest
RUN apt-get install git && cd /go/src && git clone https://github.com/koinkoin-io/koinkoin.backend && cd koinkoin.backend
CMD ["make run"]