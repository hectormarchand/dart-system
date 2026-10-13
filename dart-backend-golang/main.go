package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"dart-backend-golang/internal/darthit"
	"dart-backend-golang/internal/game"
	"dart-backend-golang/internal/gamebusiness"
	"dart-backend-golang/internal/mongodb"
	"dart-backend-golang/internal/stream"
)

func main() {
	ctx := context.Background()

	client, db, err := mongodb.Connect(ctx)
	if err != nil {
		log.Fatalf("Cannot connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	gameBusinessService := gamebusiness.New()
	gameService := game.NewService(db, gameBusinessService)
	dartHitService := darthit.NewService(gameService, gameBusinessService)
	streamService := stream.NewService()

	mux := http.NewServeMux()
	game.RegisterRoutes(mux, gameService)
	darthit.RegisterRoutes(mux, dartHitService)
	stream.RegisterRoutes(mux, streamService)

	port := os.Getenv("port")
	if port == "" {
		port = "3000"
	}

	log.Printf("🎯 Server started on port %s !", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
