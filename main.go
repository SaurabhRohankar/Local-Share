package main

import (
	"context"
	"local_share/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var serverAddress = ":8080"

func main() {
	// logger
	l := log.New(os.Stdout, "local-share", log.LstdFlags)

	// Create new ServeMux
	sm := http.NewServeMux()

	// Create handlers
	hh := handlers.NewHome(l)
	uh := handlers.NewUpload(l)

	sm.Handle("/", hh)
	sm.Handle("/upload", uh)

	// Create a new server
	s := http.Server{
		Addr:         serverAddress,
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the server in separate go routine
	go func() {
		l.Println("Starting the server on ", serverAddress)
		err := s.ListenAndServe()
		if err != nil {
			l.Println("Error starting the server: ", err)
			os.Exit(1)
		}
	}()

	// Wait for an interrupt or kill signal to stop the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// Creating context for graceful shutdown operation
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt a graceful shutdown
	err := s.Shutdown(ctx)
	if err != nil {
		l.Println("Error during server shutdown: ", err)
	}

	l.Println("Server Stopped!")

}
