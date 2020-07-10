package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/beatlabs/harvester"
	"github.com/beatlabs/harvester/sync"
)

type config struct {
	IndexName      sync.String `seed:"customers-v1"`
	CacheRetention sync.Int64  `seed:"43200" env:"ENV_CACHE_RETENTION_SECONDS"`
	LogLevel       sync.String `seed:"DEBUG" flag:"loglevel"`
}

func main() {
	ctx, cnl := context.WithCancel(context.Background())
	defer cnl()

	err := os.Setenv("ENV_CACHE_RETENTION_SECONDS", "86400")
	if err != nil {
		log.Fatalf("failed to set env var: %v", err)
	}

	cfg := config{}

	ich := make(chan int64, 1)
	lch := make(chan string, 1)
	go listen(ich, lch)

	cfg.CacheRetention.Listen(ich)
	cfg.LogLevel.Listen(lch)

	h, err := harvester.New(&cfg).Create()
	if err != nil {
		log.Fatalf("failed to create harvester: %v", err)
	}

	err = h.Harvest(ctx)
	if err != nil {
		log.Fatalf("failed to harvest configuration: %v", err)
	}

	log.Printf("Config : IndexName: %s, CacheRetention: %d, LogLevel: %s\n", cfg.IndexName.Get(), cfg.CacheRetention.Get(), cfg.LogLevel.Get())

	// let all config change handlers in the goroutine to catch values
	time.Sleep(100 * time.Millisecond)
}

func listen(cch <-chan int64, lch <-chan string) {
	for {
		select {
		case index := <-cch:
			log.Printf("Config changed. CacheRetention: %d", index)
		case logLevel := <-lch:
			log.Printf("Config changed. LogLevel: %s", logLevel)
		}
	}
}
