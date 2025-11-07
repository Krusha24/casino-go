// game/logic.go
package game

import (
	"casino/game/guessnumber"
	"casino/logservice"
	"casino/player"
	"casino/utils/io"
)

// Menu запускает главное меню казино и цикл выбора игр, пока у игрока есть деньги.
//
// Функция динамически формирует меню на основе переданного списка игр (gamesList),
// запрашивает выбор у пользователя и запускает выбранную игру через интерфейс IGame.
//
// Возвращает false, если игрок решил выйти (выбор 0) или его баланс иссяк.
func Menu(player *player.Player, io io.FullIOProvider, logger logservice.IFullLogger, gamesList []IGame) bool {
	logger.Info("Функция меню запустилась")

	for player.Balance > 0 {
		prompt := "Чего вы хотите?\nВыйти из казино - 0\n"

		for i, game := range gamesList {
			prompt += io.Swritef("%d: %s\n", i+1, game.Name())
		}

		maxChoice := len(gamesList)

		choice, err := io.ReadInt(prompt, 0, maxChoice)
		if err != nil {
			logger.Fatal("Ошибка при чтении баланса пользователя: %v", err)
			io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
			return false
		}

		switch {
		case choice == 0:
			logger.Info("Пользователь вышел из казино")
			io.WriteLine(player.StatsString())
			io.WriteLine("Всего доброго!")
			return false
		default:
			gameIndex := choice - 1

			if gameIndex >= 0 && gameIndex < len(gamesList) {
				chosenGame := gamesList[gameIndex]
				gameHasBalance := chosenGame.Play(player, io, logger)

				if !gameHasBalance {
					return false
				}
			}
		}
	}
	// Если цикл завершился из-за отсутствия денег.
	return false
}

// Start запускает приложение казино.
//
// Выполняет первоначальную инициализацию: создание списка доступных игр,
// запрос имени и стартового баланса у пользователя, создание объекта игрока,
// а затем запускает главное меню (Menu).
func Start(io io.FullIOProvider, logger logservice.IFullLogger) {
	// 1. Инициализация списка доступных игр
	// GamesList теперь определяется здесь, а не как глобальная переменная
	gamesList := []IGame{
		// Инициализируем экземпляр игры "Угадайка"
		guessnumber.NewGame(),
	}

	var name string
	var balance float64
	var prompt string

	prompt = "Привет! Как тебя зовут? "
	name, err := io.ReadString(prompt, true)
	if err != nil {
		logger.Fatal("Ошибка при чтении имени пользователя: %v", err)
		io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
		return
	}

	// Max = 0 в ReadFloat означает, что максимальное значение не ограничено
	prompt = "И сколько же у тебя денег? "
	balance, err = io.ReadFloat(prompt, 1, 0)
	if err != nil {
		logger.Fatal("Ошибка при чтении баланса пользователя: %v", err)
		io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
		return
	}

	player := player.NewPlayer(name, balance)

	io.Writef("Привет, %s! Ты начинаешь с балансом %.2f.\n", player.Name, player.Balance)
	Menu(&player, io, logger, gamesList)
}
