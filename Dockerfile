FROM golang:1.14.5-buster
RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

WORKDIR $GOPATH/src/
RUN mkdir microservices
WORKDIR $GOPATH/src/microservices
RUN mkdir src
RUN go mod init
WORKDIR $GOPATH/src/microservices/src
CMD ["watcher", "./main.go"]