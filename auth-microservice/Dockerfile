FROM golang:1.6

ADD . /go/src/github.com/s0lus/auth-microservice/

WORKDIR "/go/src/github.com/s0lus/auth-microservice"

RUN go build github.com/s0lus/auth-microservice

ENTRYPOINT ./auth-microservice

EXPOSE 8080
