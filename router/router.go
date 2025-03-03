package router

import (
	"github.com/gorilla/mux"
	"github.com/singh123ankit/Stocks/handler"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/stock/{id}", handler.GetStock).Methods("GET")
	router.HandleFunc("/api/stock", handler.GetAllStock).Methods("GET")
	router.HandleFunc("/api/stock/{id}", handler.UpdateStock).Methods("PUT")
	router.HandleFunc("/api/newstock", handler.CreateStock).Methods("POST")
	router.HandleFunc("/api/stock/{id}", handler.DeleteStock).Methods("DELETE")

	return router
}
