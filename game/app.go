// game/logic.go
package game

import (
	"casino/game/guessnumber"
	"casino/player"
	"casino/utils/io"
)

// Menu запускает главное меню казино и цикл выбора игр, пока у игрока есть деньги.
// Возвращает false, если игрок решил выйти.
func Menu(player *player.Player, io io.FullIOProvider) bool {
	var choice int
	var prompt string
	for player.Balance > 0 {
		prompt = "Чего вы хотите?\nВыйти из казино - 0\nПоиграть в угадайку - 1\n"
		choice, _ = io.ReadInt(prompt, 0, 1)
		switch {
		case choice == 0:
			io.WriteLine(player.StatsString())
			io.WriteLine("Всего доброго!")
			return false
		case choice == 1:
			// GuessNumber возвращает true, если игрок вышел из игры.
			// Мы игнорируем его возвращаемое значение и продолжаем цикл Menu.
			guessnumber.Play(player, io)
		}
	}
	return false
}

// Start запускает игру, запрашивая у пользователя имя и стартовый баланс,
// затем запускает главное меню.
func Start(io io.FullIOProvider) {
	var name string
	var balance float64
	var prompt string

	prompt = "Привет! Как тебя зовут? "
	name, _ = io.ReadString(prompt, true)

	// Max = 0 в ReadFloat означает, что максимальное значение не ограничено
	prompt = "И сколько же у тебя денег? "
	balance, _ = io.ReadFloat(prompt, 1, 0)

	player := player.NewPlayer(name, balance)

	io.Writef("Привет, %s! Ты начинаешь с балансом %.2f.\n", player.Name, player.Balance)
	Menu(&player, io)
}
