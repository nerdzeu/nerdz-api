FROM golang:latest
MAINTAINER Paolo Galeone <nessuno@nerdz.eu>

ADD . /go/src/github.com/nerdzeu/nerdz-api

RUN go install github.com/nerdzeu/nerdz-api
ENTRYPOINT /go/bin/nerdz-api
EXPOSE 8080
