package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Subhrakanti2025/golang-olx-api/internal/config"
	"github.com/Subhrakanti2025/golang-olx-api/internal/handlers"
)

func main() {
	// why w is not pointer and why r is pointer

	cfg := config.MustLoad()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlers.Health)
	fmt.Println("Running")
	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	// err := http.ListenAndServe(":8000", mux)
	// err := srv.ListenAndServe()
	// if err != nil {
	// 	log.Fatal("Servered faild: %v", err)
	// }
	// shorthand of above
	log.Printf("Server is listinig... %v", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Servered faild: %v", err)
	}
}
