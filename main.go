package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	simpleEndpoint = "/simple"
	logLevel       = "debug"
)

func main() {
	// Create a new logger
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		fmt.Printf("Error parsing log level: error=%+v\n", err)
		os.Exit(2)
	}
	l.SetLevel(level)

	mux := http.NewServeMux()
	listenAddr := fmt.Sprintf("%s:%s", "", "8080")
	server := &http.Server{
		Addr:              listenAddr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Create channels to listen for server errors and OS interrupts
	serverErrors := make(chan error, 1)
	osInterrupt := make(chan os.Signal, 1)

	defer close(osInterrupt)
	signal.Notify(osInterrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		l.Info(fmt.Sprintf("Sample server started on address : %s", listenAddr))
		serverErrors <- server.ListenAndServe()
	}()

	// Listen
	select {
	case err := <-serverErrors:
		l.Fatalf("error starting server: %v", err)
	case <-osInterrupt:
		fmt.Sprintln("received OS interrupt, shutting down")
		grace := time.Duration(5 * time.Second)
		ctx, cancel := context.WithTimeout(context.Background(), grace)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			l.Fatalf("Gracefule shutdown of server failed: %v", err)
			if err := server.Close(); err != nil {
				l.Errorf("Couldnt close the server: %v", err)
			}
		}
		l.Info("Server shutdown completed")
	}
}
