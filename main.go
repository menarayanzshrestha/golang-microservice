package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"intro/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	bb := handlers.NewBye(l)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/bye", bb)
	sm.Handle("/product", ph)

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { //default
	// 	log.Println("he")

	// })
	// http.HandleFunc("/goodby", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("goodbye")
	// })

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Println("Running on 9090")

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
