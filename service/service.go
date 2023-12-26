package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func main() {
	cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name: "Test",
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 1 && failureRatio >= 0.6
		},
	})
	server()
}

func server() {
	router := chi.NewMux()
	router.Use(middleware.Logger)

	router.Get("/{ID}", handleReq)
	http.ListenAndServe(":8090", router)
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	resp, err := requestWithCB(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func requestWithCB(id string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		client := &http.Client{
			Timeout: 20 * time.Second,
		}
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://192.168.18.8:8080/user/%s", id), nil)
		if err != nil {
			return nil, err
		}

		response, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			fmt.Println("Requisicao falhou")
			return nil, fmt.Errorf("invalid status code")
		}

		// transformar a resposta do server com a estrutura contida nele e retonrar para o request

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}
	return body.([]byte), nil
}
