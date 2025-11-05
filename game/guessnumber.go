package game

import (
	"casino/player"
	"casino/utils"
	"casino/utils/input"
	"fmt"
)

// GuessNumber запускает цикл игры "Угадай число" для игрока.
// Возвращает true, если игрок решил выйти до потери всего баланса,
// и false, если игрок потерял весь баланс.
func GuessNumber(player *player.Player, input input.InputProvider) bool {
	var choice int
	var prompt string
	var choosenNumber int

	// Цикл продолжается, пока у игрока есть деньги
	for player.Balance > 0 {
		prompt = "Если вы захотите выйти из угадайки - нажмите 0\nЕсли хотите продолжить - нажмите 1\n"
		choice, _ = input.ReadInt(prompt, 0, 1)

		switch {
		case choice == 0:
			// Пользователь решил выйти. Возвращаем true (успешный выход)
			return true
		case choice == 1:
			fmt.Printf("Ваш баланс - %.2f\n", player.Balance)
			prompt = "Сколько правильных чисел вы хотите?\n1: коэффициент - 10\n2: коэффициент - 4.5\n3: коэффициент - 2.5\n4: коэффициент - 1.7\n"
			difficult, _ := input.ReadInt(prompt, 1, 4)
			var winIndexes = utils.CreateWinIndexes(difficult)

			prompt = "Введите вашу ставку: "
			player.Bet, _ = input.ReadFloat(prompt, 0, player.Balance)

			prompt = "Выберите выйграшное число от 1 до 10: "
			choosenNumber, _ = input.ReadInt(prompt, 1, 10)

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
