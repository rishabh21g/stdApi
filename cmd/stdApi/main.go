package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rishabh21g/stdapi/internal/config"
)

func main() {
	fmt.Println("Hello, to Student API!")

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
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start a server")
	}

	// graceful shutdown
}
