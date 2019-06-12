FROM golang:1.8

ADD . /go/src/github.com/redeam
RUN go install github.com/redeam

ENTRYPOINT /go/bin/redeam

EXPOSE 8080