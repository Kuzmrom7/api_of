FROM golang:1.9

MAINTAINER orderfood

RUN go get -u github.com/kardianos/govendor

WORKDIR /go/src/github.com/orderfood/api_of/cmd/api
ADD . /go/src/github.com/orderfood/api_of

RUN govendor sync

RUN go build

ENTRYPOINT ["/go/src/github.com/orderfood/api_of/cmd/api"]