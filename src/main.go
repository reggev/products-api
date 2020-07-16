package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "oops", http.StatusBadRequest)
		} else {
			fmt.Fprintf(w, "hello %s", payload)
		}
	})
	http.HandleFunc("/going-home", func(w http.ResponseWriter, r *http.Request) {
		log.Println("going-home")
	})

	// :9090 -> bind to any address on my machine on port 9090
	http.ListenAndServe(":9090", nil)
}
