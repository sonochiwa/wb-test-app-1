package utils

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

// RemoveDuplicates - функция, которая убирает дубликаты из массива
func RemoveDuplicates(numbers []int) []int {
	encountered := map[int]bool{}
	var result []int

	for v := range numbers {
		if encountered[numbers[v]] == true {
		} else {
			encountered[numbers[v]] = true
			result = append(result, numbers[v])
		}
	}

	return result
}

// ConvertIntArrToString - функция, которая преобразует массив чисел в строку
func ConvertIntArrToString(numbers []int) string {
	// Сортируем массив чисел, чтобы для одинакового множества независимо от
	// порядка значений всегда генерировалась одинаковая строка
	slices.Sort(numbers)

	// Преобразуем []int в string
	result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), " "), "")

	return result
}

// GetResult - функция, которая может выполняться достаточно долгое время
func GetResult(x int64) (int64, error) {
	// Ожидание случайного времени (от 1 до 6 секунд)

	// Имитируем выполнение тяжелой задачи длительностью от 1 до 5 секунд
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

	// Выход = вход * вход, ошибка пустая
	return x * x, nil
}
