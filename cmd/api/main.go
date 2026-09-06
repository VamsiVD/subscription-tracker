package main

import (
	"context"
	"log"
	"subscriptionTracker/internal/config"
	"subscriptionTracker/internal/db"
	"subscriptionTracker/internal/server"
	"subscriptionTracker/internal/subscriptions"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := db.Connect(context.Background(), cfg.DatabaseUrl)

	if err != nil {
		log.Fatalf("error connect to database: %v", err)
	}

	defer db.Close()

	repo := subscriptions.NewRespository(db)
	handler := subscriptions.NewHandler(repo)
	r := server.NewRouter(handler)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
