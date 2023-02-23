package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// StartApplication is the entry point of the application
func StartApplication() {
	// ...
	log.Println("Starting server...")

	// I would initialize database connection here handler := handlers.NewHandler(db)
	// injection other services
	router, err := inject()

	if err != nil {
		log.Fatalf("Failure to inject data sources: %v\n", err)
	}

	// grabing port from env for running server local or on heroku
	port := os.Getenv("API_PORT")

	// if port env is empty then make it default 8080
	if port == "" {
		port = "50052"
	}

	// Graceful server shutdown
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		// running the server on localhost with given port
		// ------------------------ This CODE HERE is for TLS --------------------------
		// if err := srv.ListenAndServeTLS("./server.crt", "./server.pem"); err != nil && err != http.ErrServerClosed {
		// 	log.Fatalf("Failed to intialized server: %v\n", err)
		// }
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to intialized server: %v\n", err)
		}
	}()

	fmt.Printf("Listining on port %s\n", port)

	// wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
