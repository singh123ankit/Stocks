package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	db "github.com/singh123ankit/Stocks/common/postgresqldriver"
	"github.com/singh123ankit/Stocks/models"
)

type response struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, "Error in decoding request body", http.StatusBadRequest)
		return
	}
	insertID, err := db.InsertStock(stock)
	if err != nil {
		http.Error(w, "Error while inserting record", http.StatusInternalServerError)
		return
	}
	res := response{
		ID:      insertID,
		Message: "Stock created successfully",
	}
	json.NewEncoder(w).Encode(res)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error in converting id to integer", http.StatusInternalServerError)
		return
	}
	stock, err = db.GetStockById(int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Failed to fetch record by id.")
			http.Error(w, "No record found for this ID", http.StatusBadRequest)
			return
		}
		log.Printf("Unable to get record: %v", err)
		http.Error(w, "Failed to connect to the database !", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(stock)
	if err != nil {
		http.Error(w, "Error in encoding requested data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := db.GetAllStocks()
	if err != nil {
		http.Error(w, "Error in fetching requested data", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(stocks)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, "Error in decoding request body", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error in converting id to integer", http.StatusInternalServerError)
		return
	}
	updatedRows, err := db.UpdateStock(int64(id), stock)
	if err != nil {
		http.Error(w, "Error while patching data", http.StatusInternalServerError)
		return
	}
	msg := fmt.Sprintf("Stock updated successfully. Total rows/records affected %v", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error in converting id to integer", http.StatusInternalServerError)
		return
	}
	deletedRows, err := db.DeleteStock(int64(id))
	if err != nil {
		http.Error(w, "Error while deleting a record", http.StatusInternalServerError)
		return
	}
	msg := fmt.Sprintf("Stock deleted successfully. Total rows/records %v", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
}
