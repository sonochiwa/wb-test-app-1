package utils

import (
	"testing"
)

func TestConvertIntArrToString(t *testing.T) {
	// Тест с неотсортированным массивом
	t.Run("UnsortedArray", func(t *testing.T) {
		input := []int{3, 1, 4, 1, 5, 9}
		expected := "[1 1 3 4 5 9]"
		result := ConvertIntArrToString(input)
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})

	// Тест с отсортированным массивом
	t.Run("SortedArray", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := "[1 2 3 4 5]"
		result := ConvertIntArrToString(input)
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})

	// Тест с пустым массивом
	t.Run("EmptyArray", func(t *testing.T) {
		input := []int{}
		expected := "[]"
		result := ConvertIntArrToString(input)
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})

	// Тест с одинаковыми значениями в разном порядке
	t.Run("DifferentOrder", func(t *testing.T) {
		input := []int{5, 4, 3, 2, 1}
		expected := "[1 2 3 4 5]"
		result := ConvertIntArrToString(input)
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})

	// Тест с отрицательными числами
	t.Run("NegativeNumbers", func(t *testing.T) {
		input := []int{-5, 3, -1, 2, 0}
		expected := "[-5 -1 0 2 3]"
		result := ConvertIntArrToString(input)
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})
}

func TestGetResult(t *testing.T) {
	// Тест с положительным входом
	t.Run("PositiveInput", func(t *testing.T) {
		input := int64(5)
		expectedResult := int64(25)

		result, err := GetResult(input)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != expectedResult {
			t.Errorf("Expected result: %v, Got: %v", expectedResult, result)
		}
	})

	// Тест с нулевым входом
	t.Run("ZeroInput", func(t *testing.T) {
		input := int64(0)
		expectedResult := int64(0)

		result, err := GetResult(input)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != expectedResult {
			t.Errorf("Expected result: %v, Got: %v", expectedResult, result)
		}
	})
}
