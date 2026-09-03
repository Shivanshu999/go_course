package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/students-api/internal/config"
)

func main() {

	//load config
	cfg := config.MustLoad()

	fmt.Println("Welcome to students api")

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	fmt.Printf("server started at %s\n", cfg.Address)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}

}
