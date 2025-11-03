// player/player.go
package player

import "fmt"

// Player представляет игрока казино, его текущий баланс и статистику.
type Player struct {
	Name        string
	Balance     float64
	Bet         float64
	Wins        int
	Loses       int
	TotalProfit float64
}

// NewPlayer создает новый экземпляр игрока с указанным именем и начальным балансом.
func NewPlayer(name string, balance float64) Player {
	return Player{Name: name, Balance: balance}
}

// Win обновляет баланс игрока при выигрыше.
// Обнуляет текущую ставку (p.Bet), увеличивает счетчик побед и общий выигрыш.
func (p *Player) Win(amount float64) {
	p.Balance += amount
	p.Bet = 0
	p.Wins++
	p.TotalProfit += amount
}

// Lose обновляет баланс игрока при проигрыше.
// Вычитает текущую ставку (p.Bet) из баланса, обнуляет ставку,
// увеличивает счетчик проигрышей и уменьшает общий выигрыш.
func (p *Player) Lose() {
	p.Balance -= p.Bet
	p.TotalProfit -= p.Bet
	p.Bet = 0
	p.Loses++

}

// StatsString возвращает строку, содержащую текущую статистику
// игрока в удобном для чтения формате.
func (p *Player) StatsString() string {
	return fmt.Sprintf("Победы: %d, Проигрыши: %d, Тотал: %.2f", p.Wins, p.Loses, p.TotalProfit)
}
