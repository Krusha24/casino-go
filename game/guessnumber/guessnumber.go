package guessnumber

import (
	"casino/logservice"
	"casino/player"
	"casino/utils"
	"casino/utils/io"
)

// GuessNumber запускает цикл игры "Угадай число" для игрока.
// Возвращает true, если игрок решил выйти до потери всего баланса,
// и false, если игрок потерял весь баланс.
func Play(player *player.Player, io io.FullIOProvider, logger logservice.IFullLogger) bool {
	logger.Info("Пользователь зашел в угадайку")
	var prompt string
	var choosenNumber int

	// Цикл продолжается, пока у игрока есть деньги
	for player.Balance > 0 {
		prompt = "Если вы захотите выйти из угадайки - нажмите 0\nЕсли хотите продолжить - нажмите 1\n"
		choice, err := io.ReadInt(prompt, 0, 1)
		if err != nil {
			logger.Fatal("Ошибка при чтении ввода пользователя: %v", err)
			io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
			return false
		}

		switch {
		case choice == 0:
			// Пользователь решил выйти. Возвращаем true (успешный выход)
			return true
		case choice == 1:
			io.Writef("Ваш баланс - %.2f\n", player.Balance)
			prompt = "Сколько правильных чисел вы хотите?\n1: коэффициент - 10\n2: коэффициент - 4.5\n3: коэффициент - 2.5\n4: коэффициент - 1.7\n"
			difficult, err := io.ReadInt(prompt, 1, 4)
			if err != nil {
				logger.Fatal("Ошибка при чтении ввода пользователя: %v", err)
				io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
				return false
			}

			var winIndexes = utils.CreateWinIndexes(difficult)

			prompt = "Введите вашу ставку: "
			player.Bet, err = io.ReadFloat(prompt, 0, player.Balance)
			if err != nil {
				logger.Fatal("Ошибка при чтении ввода пользователя: %v", err)
				io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
				return false
			}

			prompt = "Выберите выйграшное число от 1 до 10: "
			choosenNumber, _ = io.ReadInt(prompt, 1, 10)

			isWinner := false
			for _, value := range winIndexes {
				if value == choosenNumber {
					winningSum := calculateWinnings(difficult, player.Bet)
					isWinner = true
					bet := player.Bet
					player.Win(winningSum)
					logger.LogGame(player.Name, bet, "Win", player.Balance)

					io.Writef("Вы угадали! Ваш баланс - %.2f\n", player.Balance)

					break
				}
			}

			if !isWinner {
				bet := player.Bet
				player.Lose()
				logger.LogGame(player.Name, bet, "Lose", player.Balance)
				io.Writef("Вы проиграли, ваш баланс - %.2f\n", player.Balance)
			}
		}
	}
	io.WriteLine("У вас закончились деньги.")
	io.WriteLine(player.StatsString())
	return false // Возвращаем false (вынужденный выход/проигрыш)
}
