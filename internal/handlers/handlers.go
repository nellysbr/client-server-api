package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/nellysbr/client-server-api/internal/database"
	"github.com/nellysbr/client-server-api/internal/models"
)

func GetQuotationHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
		defer cancel()

		quotation, err := fetchQuotation(ctx)
		if err != nil {
			http.Error(w, "Falha em realizar o fetch", http.StatusInternalServerError)
			return
		}

		bid, err := strconv.ParseFloat(quotation.Bid, 64)
		if err != nil {
			http.Error(w, "Falha ao converter para float64", http.StatusInternalServerError)
			return
		}
		
		err = database.SaveQuotation(ctx, db, bid)
		if err != nil {
			http.Error(w, "Falha ao salvar a quotação", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(quotation)
	}
}

func fetchQuotation(ctx context.Context) (*models.Quotation, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]models.Quotation
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	quotation := result["USDBRL"]
	return &quotation, nil
}