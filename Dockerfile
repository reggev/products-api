FROM golang:1.14.5-buster
RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

WORKDIR $GOPATH/src/
CMD ["watcher", "./main.go"]