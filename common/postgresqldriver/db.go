package postgresqldriver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/singh123ankit/Stocks/models"
)

const (
	DBDRIVER = "postgres"
)

var dbH *sql.DB
var once sync.Once

// Singleton Design Pattern
func InitDB() *sql.DB {
	var err error

	once.Do(func() {
		connStr, ok := os.LookupEnv("DATABASE_URL")
		if !ok {
			log.Println("Environment variable DATABASE_URL is not set.")
		}
		for i := 1; i <= 10; i++ {
			dbH, err = sql.Open(DBDRIVER, connStr)
			if err != nil {
				log.Printf("Error: %v | Retrying to connect to the database after 10s......", err)
				time.Sleep(10 * time.Second)
				continue
			}
			dbH.SetMaxOpenConns(10)
			dbH.SetMaxIdleConns(5)
			dbH.SetConnMaxLifetime(30 * time.Minute)
			break
		}
		log.Printf("Database connection handler dbH: %v\n", dbH)
	})
	return dbH
}

func InsertStock(stock models.Stock) (int64, error) {
	stmt := `INSERT INTO stocks(name,price,company) VALUES($1, $2, $3) RETURNING stockid`
	var id int64
	err := dbH.QueryRow(stmt, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Printf("Failed to insert Stock record into the database: %v", err)
		return id, err
	}
	fmt.Printf("Inserted record successfully %v", id)
	return id, nil
}

func GetStockById(id int64) (models.Stock, error) {
	stmt := `SELECT * FROM stocks WHERE stockid = $1`
	var stock models.Stock
	row := dbH.QueryRow(stmt, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	return stock, err
}

func GetAllStocks() ([]models.Stock, error) {
	var stocks []models.Stock
	stmt := `SELECT * FROM stocks`
	rows, err := dbH.Query(stmt)
	if err != nil {
		log.Printf("Unable to execute the query: %v", err)
		return stocks, err
	}

	defer rows.Close()
	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Printf("Unable to scan the row %v", err)
			return stocks, err
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}

func UpdateStock(id int64, stock models.Stock) (int64, error) {
	var rowsAffected int64

	stmt := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`
	res, err := dbH.Exec(stmt, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Printf("Unable to execute the query %v", err)
		return rowsAffected, err
	}
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error while checking the affected rows %v", err)
		return rowsAffected, err
	}
	fmt.Printf("Total rows/records affected: %v\n", rowsAffected)
	return rowsAffected, nil
}

func DeleteStock(id int64) (int64, error) {
	var rowsAffected int64

	stmt := `DELETE FROM stocks WHERE stockid=$1`
	res, err := dbH.Exec(stmt, id)
	if err != nil {
		log.Printf("Error while executing query %v", err)
		return rowsAffected, err
	}
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error while checking the affected rows %v", err)
		return rowsAffected, err
	}
	fmt.Printf("Total rows/records affected: %v", rowsAffected)
	return rowsAffected, nil
}
