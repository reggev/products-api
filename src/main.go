package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"products-api/src/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	productsHandler := handlers.NewProducts(logger)

	serverMux := mux.NewRouter()

	getRouter := serverMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productsHandler.GetProducts)

	putRouter := serverMux.Methods(http.MethodPut).Subrouter()
	putRouter.Use(handlers.ProductValidationMiddleware)
	putRouter.HandleFunc("/{id:[0-9]+}", productsHandler.UpdateProduct)

	postRouter := serverMux.Methods(http.MethodPost).Subrouter()
	postRouter.Use(handlers.ProductValidationMiddleware)
	postRouter.HandleFunc("/", productsHandler.AddProduct)

	server := &http.Server{
		// :9090 -> bind to any address on my machine on port 9090
		Addr:         ":9090",
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		// this is a blocking method, this is why it runs on a coroutine
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	// this is a blocking action,
	// the code is blocked until a message is fed into the channel
	sig := <-sigChannel
	logger.Println("received terminate signal, gracefull shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
