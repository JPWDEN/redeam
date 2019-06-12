FROM golang:1.11    #golang base image

WORKDIR $GOPATH/src/github.com/JPWDEN/redeam

COPY . .

EXPOSE 8080
