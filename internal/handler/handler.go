package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sonochiwa/wb-test-app/internal/global"
	"github.com/sonochiwa/wb-test-app/internal/schemas"
	"github.com/sonochiwa/wb-test-app/internal/service"
	"github.com/sonochiwa/wb-test-app/internal/utils"
)

func InitRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/process-numbers", sendNumbersSetHandler)

	return r
}

func sendNumbersSetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request schemas.NumbersSetRequestSchema

	// Обрабатываем параметр timeout из HTTP запроса
	timeout := time.Duration(2) // По дефолту 2
	if r.URL.Query().Get("timeout") != "" {
		val, err := strconv.Atoi(r.URL.Query().Get("timeout"))
		if err != nil {
			log.Println(err)
			return
		}
		timeout = time.Duration(val)
	}

	// Парсим тело запроса и проверяем на валидность
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Преобразуем массив чисел в строку для использования в качестве ключа в хранилище
	numbers := utils.ConvertIntArrToString(request.Numbers)

	// Если в хранилище уже есть результат для ключа numbers, то возвращаем этот результат клиенту
	// иначе отправляемся обрабатывать данные запроса
	if value, ok := global.Storage[numbers]; ok {
		if err := json.NewEncoder(w).Encode(value); err != nil {
			log.Println(err)
		}
		return
	}

	done := make(chan struct{})
	go func() {
		// Обрабатываем данные запроса
		service.ProcessNumbers(request, numbers)

		// После успешной обработки сигнализируем об этом в канал done
		done <- struct{}{}
	}()

	var response any // Ответ, который отправим клиенту

	select {
	case <-done: // Обработка данных завершена успешно
		response = global.Storage[numbers]
		break
	case <-time.After(timeout * time.Second): // Если время вышло, и нужно давать клиенту ответ
		response = map[string]string{"message": "Не все данные успели обработаться, попробуйте запросить их позже"}
		break
	}

	// Формируем HTTP-ответ
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
	}
}
