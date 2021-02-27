FROM golang:1.12

ADD . /go/src/github.com/matzhouse/go-grpc/server

RUN go install github.com/matzhouse/go-grpc/server

ENTRYPOINT ["/go/bin/server"]

EXPOSE 5300
