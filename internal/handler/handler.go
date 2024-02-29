package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sonochiwa/wb-test-app/internal/schemas"
	"github.com/sonochiwa/wb-test-app/internal/service"
)

func InitRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/process-numbers", sendNumbersSetHandler)

	return r
}

func sendNumbersSetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request schemas.NumbersSetRequest
	responseCh := make(chan schemas.NumbersSetResponse)

	// Обрабатываем параметр timeout из HTTP запроса
	timeout := time.Duration(2)
	if r.URL.Query().Get("timeout") != "" {
		val, err := strconv.Atoi(r.URL.Query().Get("timeout"))
		if err != nil {
			log.Println(err)
			return
		}
		timeout = time.Duration(val)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	// Парсим тело запроса и проверяем на валидность
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обрабатываем данные запроса
	go service.ProcessNumbers(ctx, request, responseCh)

	// Формируем HTTP-ответ
	if err := json.NewEncoder(w).Encode(<-responseCh); err != nil {
		log.Println(err)
	}
}
