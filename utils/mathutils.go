// utils/mathutils.go
package utils

import (
	"math/rand"
)

// randRange генерирует случайное целое
// число в заданном диапазоне [min, max] включительно.
func randRange(min, max int) int {
	// rand.Intn(n) возвращает число в диапазоне [0, n).
	// Мы используем max+1-min, чтобы включить max.
	return rand.Intn(max+1-min) + min
}

// CreateWinIndexes генерирует слайс уникальных выигрышных чисел на основе
// заданной сложности (difficult).
//
// Числа генерируются в диапазоне от 1 до 10. Количество уникальных чисел
// равно значению difficult. Возвращает слайс целых чисел.
func CreateWinIndexes(difficult int) []int {
	// Используем map для эффективного отслеживания уникальности (O(1) поиск).
	// Значение true просто отмечает присутствие ключа.
	uniqueNumbers := make(map[int]bool)

	// Цикл повторяется, пока мы не сгенерируем нужное количество уникальных чисел
	for len(uniqueNumbers) < difficult {
		// Генерируем число от 1 до 10 включительно.
		randomValue := randRange(1, 10)

		// Проверка: находится ли число уже в нашем наборе?
		if !uniqueNumbers[randomValue] {
			// Если нет, добавляем его.
			uniqueNumbers[randomValue] = true
		}
	}

	// Преобразуем map обратно в слайс для возврата.
	winIndexesSlice := make([]int, 0, difficult)
	for num := range uniqueNumbers {
		winIndexesSlice = append(winIndexesSlice, num)
	}
	return winIndexesSlice
}

// CalculateWinnings рассчитывает сумму выигрыша
// исходя из сложности (difficult) и ставки (bet).
//
// Использует предопределенную таблицу коэффициентов.
func CalculateWinnings(difficult int, bet float64) float64 {
	x := map[int]float64{1: 10, 2: 4.5, 3: 2.5, 4: 1.7}

	// Сумма выигрыша: ставка * коэффициент
	winningAmount := bet*x[difficult] - bet
	return winningAmount
}
