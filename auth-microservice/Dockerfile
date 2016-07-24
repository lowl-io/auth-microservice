FROM golang:1.6

ADD . /go/src/github.com/s0lus/auth-microservice

RUN go install github.com/s0lus/auth-microservice

ENTRYPOINT /go/bin/auth-microservice

EXPOSE 8080
