FROM golang:1.14.5-buster
RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

WORKDIR $GOPATH/src/
RUN mkdir products-api
WORKDIR $GOPATH/src/products-api
RUN mkdir src
RUN go mod init
WORKDIR $GOPATH/src/products-api/src
CMD ["watcher", "./main.go"]