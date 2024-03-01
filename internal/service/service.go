package service

import (
	"context"
	"log"
	"strconv"

	"github.com/sonochiwa/wb-test-app/internal/global"
	"github.com/sonochiwa/wb-test-app/internal/schemas"
	"github.com/sonochiwa/wb-test-app/internal/utils"
)

func ProcessNumbers(ctx context.Context, request schemas.NumbersSetRequest, responseCh chan<- schemas.NumbersSetResponse) {
	var response schemas.NumbersSetResponse
	resultsCh := make(chan map[string]int)

	numbersSet := utils.RemoveDuplicates(request.Numbers)
	numbers := utils.ConvertIntArrToString(numbersSet)

	// Если в хранилище уже есть результат, то возвращаем этот результат клиенту
	if data, ok := global.Storage.Get(numbers); ok {
		if len(data.Results) == len(numbersSet) {
			responseCh <- schemas.NumbersSetResponse{Results: data.Results}
		}
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
		for range numbersSet {
			for k, v := range <-resultsCh {
				data, _ := global.Storage.Get(numbers)
				data.Results[k] = v
				global.Storage.Set(numbers, data)
			}
		}
		data, _ := global.Storage.Get(numbers)
		responseCh <- data
	}()

	<-ctx.Done()
	data, _ := global.Storage.Get(numbers)
	response = schemas.NumbersSetResponse{
		Results: data.Results,
	}

	if len(data.Results) < len(numbersSet) {
		response.Details = "Не все данные успели обработаться, попробуйте запросить их позже"
	}

	responseCh <- response
}
