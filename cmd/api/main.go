package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Subhrakanti2025/golang-olx-api/internal/config"
	"github.com/Subhrakanti2025/golang-olx-api/internal/db"
	"github.com/Subhrakanti2025/golang-olx-api/internal/handlers"
	"github.com/Subhrakanti2025/golang-olx-api/internal/middleware"
)

func main() {
	// why w is not pointer and why r is pointer

	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("DB is not connected: %v", err)
	}
	fmt.Println("Database connected")
	loggerHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})
	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)
	mux := http.NewServeMux()
	lh := handlers.NewDbConstructor(db, logger)
	usr := handlers.NewDbConstructorForAuth(db, logger)
	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("GET /listing", lh.List)
	mux.HandleFunc("DELETE /listing/{id}", lh.DeleteList)
	mux.HandleFunc("POST /listing", lh.CreateList)
	mux.HandleFunc("POST /user", usr.CreateUser)
	fmt.Println("Running")
	handler := middleware.RequestId(mux)
	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
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
