package main

import (
	"context"
	"github.com/pandulaDW/go-microservice-with-grpc/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)
	pd := handlers.NewProducts(l)

	router := http.NewServeMux()
	router.HandleFunc("/products", pd.ServeHttp)

	server := &http.Server{
		Addr:         ":4000",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	log.Println("server listening to requests at port 4000...")

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	d := time.Now().Add(30 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		l.Fatal(err)
	}
}
