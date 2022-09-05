package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AlvinChiew/go-microservice-apps/handlers"
	"github.com/gorilla/mux"
)

func check_health(w http.ResponseWriter, r *http.Request) {
	log.Println("Application health check")
	response := map[string]string{
		"status":    "ok",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func serve_homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Homepage is served")
}

func main() {

	l := log.New(os.Stdout, "sample-api", log.LstdFlags)

	sm := mux.NewRouter()

	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/health", check_health)
	getR.HandleFunc("/", serve_homepage)
	getR.HandleFunc("/host", handlers.GetHostName)

	// create a new server
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)

}
