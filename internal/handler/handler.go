package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sonochiwa/wb-test-app/internal/global"
	"github.com/sonochiwa/wb-test-app/internal/schemas"
	"github.com/sonochiwa/wb-test-app/internal/service"
)

func InitRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/numbers-sets", sendNumbersSetHandler)

	return r
}

func sendNumbersSetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request schemas.NumbersSetRequestSchema

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numbers := convertIntArrToString(request.Numbers)

	// Если в хранилище уже есть результат для полученного набора чисел,
	// то мы просто вернем его и завершим выполнение функции
	if value, ok := global.Storage[numbers]; ok {
		if err := json.NewEncoder(w).Encode(value); err != nil {
			log.Println(err)
		}
		return
	}

	// Канал для сигнала, если данные успеют обработаться за отведенное время
	done := make(chan struct{})
	go func() {
		// Отправляем request на обработку
		if err := service.ProcessNumbers(request, numbers); err != nil {
			log.Println(err)
		}
		// После успешной обработки сигнализируем об этом в канал
		done <- struct{}{}
	}()

	select {
	case <-done: // Обработка данных завершена успешно
		result := global.Storage[numbers]
		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Println(err)
		}
		return
	case <-time.After(2 * time.Second): // Если время вышло, и нужно давать клиенту ответ
		if err := json.NewEncoder(w).Encode(map[string]string{
			"message": "Не все данные успели обработаться, попробуйте запросить их позже"},
		); err != nil {
			log.Println(err)
		}
	}
}

// convertIntArrToString - функция, которая преобразует массив чисел в строку
func convertIntArrToString(numbers []int) string {
	// Сортируем массив чисел, чтобы для одинакового множества независимо от
	// порядка значений всегда генерировалась одинаковая строка
	slices.Sort(numbers)

	// Преобразуем []int в string
	result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), " "), "")

	return result
}
