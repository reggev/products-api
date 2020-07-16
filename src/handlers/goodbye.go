package handlers

import (
	"log"
	"net/http"
)

/*
Goodbye is a request handler
*/
type Goodbye struct {
	logger *log.Logger
}

/*
NewGoodbye returns a new Goodbye instance
*/
func NewGoodbye(logger *log.Logger) *Goodbye {
	return &Goodbye{logger}
}

func (goodbye *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	goodbye.logger.Println("goodbye")
	w.Write([]byte("bye!"))
}
