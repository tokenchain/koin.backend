FROM golang:latest
RUN cd /home/go/pkg/koinkoin.backend
WORKDIR /home/go/pkg/koinkoin.backend
CMD ["make run"]