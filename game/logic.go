// game/logic.go
package game

import (
	"casino/player"
	"casino/utils"
	"fmt"
)

// GuessNumber запускает цикл игры "Угадай число" для игрока.
// Возвращает true, если игрок решил выйти до потери всего баланса,
// и false, если игрок потерял весь баланс.
func GuessNumber(player *player.Player) bool {
	var choice int
	var prompt string
	var choosenNumber int

	// Цикл продолжается, пока у игрока есть деньги
	for player.Balance > 0 {
		prompt = "Если вы захотите выйти из угадайки - нажмите 0\nЕсли хотите продолжить - нажмите 1\n"
		choice, _ = utils.ReadInt(prompt, 0, 1)

		switch {
		case choice == 0:
			// Пользователь решил выйти. Возвращаем true (успешный выход)
			return true
		case choice == 1:
			fmt.Printf("Ваш баланс - %.2f\n", player.Balance)
			prompt = "Сколько правильных чисел вы хотите?\n1: коэффициент - 5\n2: коэффициент - 2.5\n3: коэффициент - 1.67\n4: коэффициент - 1.25\n5: коэффициент - 1\n6: коэффициент - 0.83\n7: коэффициент - 0.71\n8: коэффициент - 0.62\n9: коэффициент - 0.56\n10: коэффициент - 0.5\n"
			difficult, _ := utils.ReadInt(prompt, 1, 10)
			var winIndexes = utils.CreateWinIndexes(difficult)

			prompt = "Введите вашу ставку: "
			player.Bet, _ = utils.ReadFloat(prompt, 1, player.Balance)

			prompt = "Выберите выйграшное число от 1 до 10: "
			choosenNumber, _ = utils.ReadInt(prompt, 1, 10)

			fmt.Println(winIndexes, choosenNumber)
			isWinner := false
			for _, value := range winIndexes {
				if value == choosenNumber {
					winningSum := utils.CalculateWinnings(difficult, player.Bet)
					player.Win(winningSum)
					fmt.Printf("Вы угадали! Ваш баланс - %.2f\n", player.Balance)
					isWinner = true
					break
				}
			}

			if !isWinner {
				player.Lose()
				fmt.Printf("Вы проиграли, ваш баланс - %.2f\n", player.Balance)
			}
		}
	}
	fmt.Println("У вас закончились деньги.")
	fmt.Println(player.StatsString())
	return false // Возвращаем false (вынужденный выход/проигрыш)
}

// Menu запускает главное меню казино и цикл выбора игр, пока у игрока есть деньги.
// Возвращает false, если игрок решил выйти.
func Menu(player *player.Player) bool {
	var choice int
	var prompt string
	for player.Balance > 0 {
		prompt = "Чего вы хотите?\nВыйти из казино - 0\nПоиграть в угадайку - 1\n"
		choice, _ = utils.ReadInt(prompt, 0, 1)
		switch {
		case choice == 0:
			fmt.Println(player.StatsString())
			fmt.Println("Всего доброго!")
			return false
		case choice == 1:
			// GuessNumber возвращает true, если игрок вышел из игры.
			// Мы игнорируем его возвращаемое значение и продолжаем цикл Menu.
			GuessNumber(player)
		}
	}
	return false
}

// Start запускает игру, запрашивая у пользователя имя и стартовый баланс,
// затем запускает главное меню.
func Start() {
	var name string
	var balance float64
	var prompt string

	prompt = "Привет! Как тебя зовут? "
	name, _ = utils.ReadString(prompt, true)

	// Max = 0 в ReadFloat означает, что максимальное значение не ограничено
	prompt = "И сколько же у тебя денег? "
	balance, _ = utils.ReadFloat(prompt, 1, 0)

	player := player.NewPlayer(name, balance)

	fmt.Printf("Привет, %s! Ты начинаешь с балансом %.2f.\n", player.Name, player.Balance)
	Menu(&player)
}
