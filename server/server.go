package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/johnsoncwb/api-test/server/models"

	"github.com/go-chi/chi/v5"
)

var counter int

func main() {
	server()
}

func server() {
	router := chi.NewMux()
	router.Use(middleware.Logger)

	router.Get("/", handleReq)
	router.Get("/user/{ID}", getUser)
	http.ListenAndServe(":8080", router)
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TUDO CERTO"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	data, err := readFile()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error to read file: %s", err)))
	}

	var users []models.User

	err = json.Unmarshal(data, &users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("erro ao efetuar o unmarshal dos dados"))
		return
	}

	id := chi.URLParam(r, "ID")

	intID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("erro ao transformar id para inteiro"))
		return
	}

	for _, user := range users {
		if user.ID == intID && counterIsInvalid() {
			data, err := json.Marshal(user)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("erro ao preparar payload para resposta"))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		}
		counter++

		if !counterIsInvalid() {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	return
}

func readFile() ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fileName := "users.json"
	return os.ReadFile(fmt.Sprintf("%s/server/%s", dir, fileName))
}

func counterIsInvalid() bool {
	return counter < 20 || (counter < 80 && counter > 79)
}
