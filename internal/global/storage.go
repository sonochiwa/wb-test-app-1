package global

import "github.com/sonochiwa/wb-test-app/internal/schemas"

// Storage - объявление глобальной переменной для хранения данных,
// чтобы избежать необходимости добавления работы с БД в тестовом задании
var Storage map[string]schemas.NumbersSetResponseSchema

func init() {
	// Инициализация хранилища
	Storage = make(map[string]schemas.NumbersSetResponseSchema)
}
