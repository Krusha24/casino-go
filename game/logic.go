// game/logic.go
package game

import (
	"casino/player"
	"casino/utils/input"
	"fmt"
)

// Menu запускает главное меню казино и цикл выбора игр, пока у игрока есть деньги.
// Возвращает false, если игрок решил выйти.
func Menu(player *player.Player, input input.InputProvider) bool {
	var choice int
	var prompt string
	for player.Balance > 0 {
		prompt = "Чего вы хотите?\nВыйти из казино - 0\nПоиграть в угадайку - 1\n"
		choice, _ = input.ReadInt(prompt, 0, 1)
		switch {
		case choice == 0:
			fmt.Println(player.StatsString())
			fmt.Println("Всего доброго!")
			return false
		case choice == 1:
			// GuessNumber возвращает true, если игрок вышел из игры.
			// Мы игнорируем его возвращаемое значение и продолжаем цикл Menu.
			GuessNumber(player, input)
		}
	}
	return false
}

// Start запускает игру, запрашивая у пользователя имя и стартовый баланс,
// затем запускает главное меню.
func Start(input input.InputProvider) {
	var name string
	var balance float64
	var prompt string

	prompt = "Привет! Как тебя зовут? "
	name, _ = input.ReadString(prompt, true)

	// Max = 0 в ReadFloat означает, что максимальное значение не ограничено
	prompt = "И сколько же у тебя денег? "
	balance, _ = input.ReadFloat(prompt, 1, 0)

	player := player.NewPlayer(name, balance)

	fmt.Printf("Привет, %s! Ты начинаешь с балансом %.2f.\n", player.Name, player.Balance)
	Menu(&player, input)
}
