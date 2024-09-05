package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nellysbr/client-server-api/internal/database"
	"github.com/nellysbr/client-server-api/internal/handlers"
)

func main() {
	db, err := sql.Open("sqlite3", "./quotations.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = database.InitDB(db)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/cotacao", handlers.GetQuotationHandler(db))

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}