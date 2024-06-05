package main

import (
	"github.com/evanstukalov/wildberries_internship_l0/internal/api"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/evanstukalov/wildberries_internship_l0/internal/cache"
	"github.com/evanstukalov/wildberries_internship_l0/internal/consumer"
	"github.com/evanstukalov/wildberries_internship_l0/internal/database"
	"github.com/evanstukalov/wildberries_internship_l0/internal/services"
)

func main() {

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	inMemoryCache := cache.NewInMemoryCache()
	orderService := services.NewMessageService(inMemoryCache, db)

	err = orderService.FillCacheWithOrders()
	if err != nil {
		return
	}

	natsConsumer, err := consumer.NewConsumer("nats://localhost:4222", orderService)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	if err := natsConsumer.Consume("orders"); err != nil {
		log.Fatalf("Failed to subscribe to NATS: %v", err)
	}

	go api.NewServer(inMemoryCache).StartHTTPServer()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Waiting for messages. Press Ctrl+C to exit.")
	sig := <-sigs
	log.Printf("Received signal %s, shutting down...", sig)
}
