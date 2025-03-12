package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"context"
	"os"
	"os/signal"

	psq "github.com/singh123ankit/Stocks/common/postgresqldriver"
	"github.com/singh123ankit/Stocks/router"
)

func main() {
	dbH := psq.InitDB()
	r := router.Router()
	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	go startServer(server)
	fmt.Println("Starting Server at 8000:")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan
	fmt.Println("Received Interrupt signal", sig)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer func() {
		if dbH != nil {
			dbH.Close()
		}
		cancel()
	}()
	server.Shutdown(ctx)
}

func startServer(server *http.Server) {
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
