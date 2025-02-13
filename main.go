package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	psq "github.com/singh123ankit/Stocks/common/postgresqldriver"
	"github.com/singh123ankit/Stocks/router"
)

var gExitChan = make(chan error)

func main() {
	dbH := psq.InitDB()
	defer func() {
		if dbH != nil {
			dbH.Close()
		}
	}()
	r := router.Router()
	startServer(r)
	fmt.Println("Starting Server at 8000:")
	err := <-gExitChan
	log.Fatal(err)
}

func startServer(r *mux.Router) {
	go func(r *mux.Router) {
		err := http.ListenAndServe(":8000", r)
		if err != nil {
			gExitChan <- err
		}
	}(r)
}
