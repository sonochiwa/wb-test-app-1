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
				log.Println("Не удалось вычислить значение для числа ", num)
				return
			}

			mutex.Lock()
			results[strconv.Itoa(num)] = int(result) // Записываем результат в общий словарь
			mutex.Unlock()
		}(num)
	}

	// Ожидаем завершения всех goroutine
	wg.Wait()

	// Записываем результат в хранилище
	mutex.Lock()
	global.Storage[numbers] = schemas.NumbersSetResponseSchema{Results: results}
	mutex.Unlock()
}
