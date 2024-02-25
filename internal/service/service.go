package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/sonochiwa/wb-test-app/internal/global"
	"github.com/sonochiwa/wb-test-app/internal/schemas"
)

func ProcessNumbers(request schemas.NumbersSetRequestSchema, numbers string) error {
	var results = make(map[string]int)
	var wg sync.WaitGroup
	mutex := sync.Mutex{}

	for _, num := range request.Numbers {
		wg.Add(1)

		go func(num int) {
			defer wg.Done()

			result, err := getResult(int64(num))

			mutex.Lock()
			defer mutex.Unlock()

			if err == nil {
				results[strconv.Itoa(num)] = int(result)
			} else {
				fmt.Printf("Error calculating result for %v: %v\n", num, err)
			}
		}(num)
	}

	wg.Wait() // Ожидаем завершения всех горутин
	global.Storage[numbers] = schemas.NumbersSetResponseSchema{Results: results}

	return nil
}

// getResult - функция, которая может выполняться достаточно долгое время
func getResult(x int64) (int64, error) {
	// Ожидание случайного времени (от 2 до 12 секунд)
	time.Sleep(time.Duration(rand.Intn(2)+2) * time.Second)

	// Выход = вход * вход, ошибка пустая
	return x * x, nil
}
