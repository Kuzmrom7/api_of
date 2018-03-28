FROM golang:1.10

MAINTAINER orderfood

RUN go get -u github.com/kardianos/govendor

WORKDIR /go/src/github.com/orderfood/api_of
ADD . .

RUN govendor sync


WORKDIR /go/src/github.com/orderfood/api_of/cmd/api

RUN chmod +x /go/src/github.com/orderfood/api_of/cmd/api/api.go

RUN go build

ENTRYPOINT ["/go/src/github.com/orderfood/api_of"]