package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
Hello is a struct implementing the handler interface
*/
type Hello struct {
	logger *log.Logger
}

/*
NewHello create a new Hello instance
*/
func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (hello *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hello.logger.Println("hello")
	payload, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "hello %s", payload)
}
