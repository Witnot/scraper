package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/Witnot/scraper/internal/api"
	"github.com/Witnot/scraper/internal/db"
	"github.com/Witnot/scraper/internal/scraper"
)

func main() {
	// Initialize DB (Init() already panics on failure)
	db.Init()

	// start REST API in a goroutine
	go api.Run() // implement Run() to start Gin on :8080

	// scheduler
	c := cron.New()
	c.AddFunc("@every 30s", func() {
		ctx := context.Background()
		targets := []string{
			"https://fakestoreapi.com/products/1",
			"https://fakestoreapi.com/products/2",
			"https://fakestoreapi.com/products/3",
			"https://fakestoreapi.com/products/4",
			"https://fakestoreapi.com/products/5",
		}

		for _, u := range targets {
			go scraper.ScrapeFakeStoreProduct(ctx, u)
		}
	})

	c.Start()

	// wait for shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	c.Stop()
	time.Sleep(1 * time.Second)
}
