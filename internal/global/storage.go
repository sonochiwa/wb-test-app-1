package global

import (
	"sync"

	"github.com/sonochiwa/wb-test-app/internal/schemas"
)

type Store struct {
	data map[string]schemas.NumbersSetResponse
	mu   sync.RWMutex
}

// NewStore - конструктор для store
func NewStore() *Store {
	return &Store{
		data: make(map[string]schemas.NumbersSetResponse),
	}
}

func (s *Store) Get(key string) (value schemas.NumbersSetResponse, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok = s.data[key]
	if !ok {
		return schemas.NumbersSetResponse{Results: map[string]int{}}, ok
	}
	return value, ok
}

func (s *Store) Set(key string, value schemas.NumbersSetResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Storage - замена базы данных
var Storage *Store

func init() {
	// Инициализация хранилища
	Storage = NewStore()
}
