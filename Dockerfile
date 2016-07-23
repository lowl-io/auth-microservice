FROM golang:1.6

EXPOSE 8080

CMD cd auth-microservice/ && go build .

ENTRYPOINT auth-microservice/auth-microservice

