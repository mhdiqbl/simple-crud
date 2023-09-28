package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	router "ottoDigital/route"
	"syscall"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	osChan := make(chan os.Signal)
	signal.Notify(osChan, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)

	server := &http.Server{
		Addr:        ":8080",
		IdleTimeout: 5 * time.Second,
		Handler:     router.Init(),
	}

	go func() {
		server.ListenAndServe()
	}()

	log.Println("server listen on port 8080")

	<-osChan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	log.Println("Signal kill triggered")
	if err := server.Shutdown(ctx); err != nil {
		cancel()
		os.Exit(1)
	}

	log.Println("Server closed...")
}
