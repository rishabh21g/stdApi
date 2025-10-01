package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rishabh21g/stdapi/internal/config"
)

func main() {

	//laod config
	cfg := config.MustLoad()

	// set up db

	// set up router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Student API"))
	})

	//set up server
	server := http.Server{Addr: cfg.Addr, Handler: router}
	fmt.Printf("Server is running on port: %s", cfg.HTTPServer.Addr)

	// listen to server
	done := make(chan os.Signal, 1) //buffered channel to avoid goroutine leak
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// os.Interrupt means ctrl+c
	// syscall.SIGINT means terminal interrupt
	// syscall.SIGTERM means termination signal
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start a server")
		}
	}()
	<-done

	// graceful shutdown
	slog.Info("Server is shutting down...")
	// server.Shutdown() we can do this but we will use context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // context in go is used to carry deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

	defer cancel() // defer keyword is used to ensure that the cancel function is called at the end of the main function execution, regardless of how the function exits (whether normally or due to an error). This is important for cleaning up resources associated with the context.

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("Server exited properly") // slog is structured logging package

}
