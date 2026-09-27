package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// why w is not pointer and why r is pointer
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})
	fmt.Println("Running")
	srv := http.Server{
		Addr:         ":8000",
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
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Servered faild: %v", err)
	}
}
