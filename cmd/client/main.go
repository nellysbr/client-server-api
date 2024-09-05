package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nellysbr/client-server-api/internal/models"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Erro criando a request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro fetching quotação:", err)
		return
	}
	defer resp.Body.Close()

	var quotation models.Quotation
	if err := json.NewDecoder(resp.Body).Decode(&quotation); err != nil {
		fmt.Println("Erro decoding:", err)
		return
	}

	err = saveQuotationToFile(quotation.Bid)
	if err != nil {
		fmt.Println("Erro salvando o arquivo:", err)
		return
	}

	fmt.Println("Quotação salva com sucesso!")
}

func saveQuotationToFile(bid string) error {
	content := fmt.Sprintf("Dólar: %s", bid)
	return os.WriteFile("cotacao.txt", []byte(content), 0644)
}