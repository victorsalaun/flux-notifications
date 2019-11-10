FROM golang:1.13 AS builder

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

EXPOSE 8080
CMD ["/go/bin/app"]