FROM golang:latest

RUN apt-get update
RUN apt-get install vim -y
RUN apt-get install -y apache2
RUN go get "github.com/go-sql-driver/mysql"

ADD . /go/src/github.com/redeam/go
RUN go install github.com/redeam/go

EXPOSE 8080
ENTRYPOINT /go/bin/redeam