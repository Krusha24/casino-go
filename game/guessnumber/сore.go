package guessnumber

// CalculateWinnings рассчитывает сумму выигрыша
// исходя из сложности (difficult) и ставки (bet).
//
// Использует предопределенную таблицу коэффициентов.
func calculateWinnings(difficult int, bet float64) float64 {
	x := map[int]float64{1: 10, 2: 4.5, 3: 2.5, 4: 1.7}

	// Сумма выигрыша: ставка * коэффициент
	winningAmount := bet*x[difficult] - bet
	return winningAmount
}