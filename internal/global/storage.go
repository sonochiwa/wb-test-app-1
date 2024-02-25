package global

import "github.com/sonochiwa/wb-test-app/internal/schemas"

// Storage - объявление глобальной переменной для хранения данных (чтобы избежать добавления БД в тестовое)
var Storage map[string]schemas.NumbersSetResponseSchema

func init() {
	// Инициализация хранилища
	Storage = make(map[string]schemas.NumbersSetResponseSchema)

	Storage["[1 2 3 4 5]"] = schemas.NumbersSetResponseSchema{
		Results: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5},
	}
}
