package service

import (
	"log"
	"strconv"
	"sync"

	"github.com/sonochiwa/wb-test-app/internal/global"
	"github.com/sonochiwa/wb-test-app/internal/schemas"
	"github.com/sonochiwa/wb-test-app/internal/utils"
)

func ProcessNumbers(request schemas.NumbersSetRequestSchema, numbers string) {
	var results = make(map[string]int)
	var wg sync.WaitGroup

	mutex := sync.Mutex{}

	// Для каждого числа из запроса создаем отдельную goroutine,
	// которая будет выполнять долгую операцию

	for _, num := range request.Numbers {
		wg.Add(1)

		go func(num int) {
			defer wg.Done()

			// Выполняем долгую операцию
			result, err := utils.GetResult(int64(num))
			if err != nil {
				mutex.Lock()
				defer mutex.Unlock()
				results[strconv.Itoa(num)] = -1 // Если произошла ошибка, записываем -1

				log.Println("Не удалось вычислить значение для числа ", num)
			}

			mutex.Lock()
			defer mutex.Unlock()
			results[strconv.Itoa(num)] = int(result) // Записываем результат в общий словарь
		}(num)
	}

	wg.Wait() // Ожидаем завершения всех goroutine

	// Записываем результат в хранилище
	global.Storage[numbers] = schemas.NumbersSetResponseSchema{Results: results}
}
