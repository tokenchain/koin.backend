FROM golang:latest

RUN mkdir -p /go/src/github.com/koin-bet/koin.backend
COPY . /go/src/github.com/koin-bet/koin.backend
WORKDIR /go/src/github.com/koin-bet/koin.backend

RUN go build -o koin *.go
RUN ls -lh

ENTRYPOINT ["koin"]
