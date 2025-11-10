package guessnumber

// CalculateNetWinnings рассчитывает сумму чистого выигрыша (Payout - Bet)
// исходя из сложности (difficult) и ставки (bet).
//
// Использует предопределенную таблицу коэффициентов.
// Коэффициенты: 1 - Самый сложный (10x), 4 - Самый легкий (1.7x).
func calculateNetWinnings(difficulty int, bet float64) float64 {
	// Коэффициенты полной выплаты (включая ставку): Payout = Bet * X[difficult]
	payoutCoefficients := map[int]float64{
		1: 10.0, // Высокий риск, высокий выигрыш
		2: 4.5,
		3: 2.5,
		4: 1.7, // Низкий риск, низкий выигрыш
	}

	// Получаем коэффициент, если он существует. Если нет, выплаты нет.
	coefficient, exists := payoutCoefficients[difficulty]
	if !exists {
		return 0.0
	}

	// Сумма полного выигрыша (ставка * коэффициент)
	totalPayout := bet * coefficient

	// Расчет чистого выигрыша (прибыли): Payout - Bet
	netWinnings := totalPayout - bet

	return netWinnings
}
