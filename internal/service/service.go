package service

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/sonochiwa/wb-test-app/internal/global"
	"github.com/sonochiwa/wb-test-app/internal/schemas"
	"github.com/sonochiwa/wb-test-app/internal/utils"
)

func ProcessNumbers(ctx context.Context, request schemas.NumbersSetRequest, responseCh chan<- schemas.NumbersSetResponse) {
	var response schemas.NumbersSetResponse
	var mu sync.RWMutex

	resultsCh := make(chan map[string]int)

	// Преобразуем массив чисел в уникальное множество
	numbersSet := utils.RemoveDuplicates(request.Numbers)

	// Преобразуем массив чисел в строку для использования в качестве ключа в хранилище
	numbers := utils.ConvertIntArrToString(numbersSet)

	// Если в хранилище уже есть результат, то возвращаем этот результат клиенту
	if value, ok := global.Storage[numbers]; ok {
		if len(global.Storage[numbers].Results) == len(numbersSet) {
			responseCh <- schemas.NumbersSetResponse{Results: value.Results}
		}
	} else {
		// Инициализируем пустой map
		mu.Lock()
		global.Storage[numbers] = schemas.NumbersSetResponse{Results: map[string]int{}}
		mu.Unlock()
	}

	// Для каждого числа из запроса создаем отдельную goroutine
	for _, num := range numbersSet {
		go func(num int) {
			result, err := utils.GetResult(int64(num)) // Выполняем долгую операцию
			if err != nil {
				log.Println("Не удалось вычислить значение для числа ", num)
				return
			}
			resultsCh <- map[string]int{strconv.Itoa(num): int(result)} // Отправляем результат в канал
		}(num)
	}

	go func() {
		// Сбор результатов из канала
		for range numbersSet {
			for k, v := range <-resultsCh {
				mu.Lock()
				global.Storage[numbers].Results[k] = v
				mu.Unlock()
			}
		}
		responseCh <- global.Storage[numbers] // Отправляем результат в канал ответа
	}()

	<-ctx.Done()
	response = schemas.NumbersSetResponse{
		Results: global.Storage[numbers].Results,
	}

	if len(response.Results) < len(numbersSet) {
		response.Details = "Не все данные успели обработаться, попробуйте запросить их позже"
	}

	responseCh <- response
}
