// utils/mathutils.go
package utils

import (
	"math/rand"
)

// randRange генерирует случайное целое число
// в заданном диапазоне [min, max] включительно.
// Эта функция приватна (начинается с маленькой буквы) и используется
// только внутри пакета utils.
func randRange(min, max int) int {
	// rand.Intn(n) возвращает число в диапазоне [0, n).
	// Мы используем max+1-min, чтобы включить max.
	return rand.Intn(max+1-min) + min
}

// CreateWinIndexes генерирует слайс указанного количества (difficult)
// уникальных случайных чисел в заданном диапазоне [1, 10] включительно.
// Количество генерируемых чисел равно значению 'difficult'.
func CreateWinIndexes(difficulty int) []int {
	// Используем map для эффективного отслеживания уникальности (O(1) поиск).
	// Значение true просто отмечает присутствие ключа.
	uniqueNumbers := make(map[int]bool)

	// Цикл повторяется, пока мы не сгенерируем нужное количество уникальных чисел
	for len(uniqueNumbers) < difficulty {
		// Генерируем число от 1 до 10 включительно.
		randomValue := randRange(1, 10)

		// Проверка: находится ли число уже в нашем наборе?
		if !uniqueNumbers[randomValue] {
			// Если нет, добавляем его.
			uniqueNumbers[randomValue] = true
		}
	}

	// Преобразуем map обратно в слайс для возврата.
	winIndexesSlice := make([]int, 0, difficulty)
	for num := range uniqueNumbers {
		winIndexesSlice = append(winIndexesSlice, num)
	}
	return winIndexesSlice
}
