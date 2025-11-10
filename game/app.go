// game/logic.go
package game

import (
	"casino/game/guessnumber"
	"casino/logservice"
	"casino/player"
	"casino/utils/ui"
	"context"
	"fmt"
)

// Menu запускает главное меню казино и цикл выбора игр, пока у игрока есть деньги.
//
// Функция динамически формирует меню на основе переданного списка игр (gamesList),
// запрашивает выбор у пользователя и запускает выбранную игру через интерфейс IGame.
//
// Возвращает false, если игрок решил выйти (выбор 0) или его баланс иссяк.
func Menu(ctx context.Context, player *player.Player, io ui.FullIOProvider, logger logservice.IFullLogger, gamesList []IGame) (bool, error) {
	logger.Info("Функция меню запустилась")

	for player.Balance > 0 {
		prompt := "Чего вы хотите?\nВыйти из казино - 0\n"

		for i, game := range gamesList {
			prompt += fmt.Sprintf("%d: %s\n", i+1, game.Name())
		}

		maxChoice := len(gamesList)

		choice, err := io.ReadIntCtx(ctx, prompt, 0, maxChoice)
		if err != nil {
			logger.Fatal("Ошибка при выборе игры пользователя: %v", err)
			io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
			return false, err
		}

		switch {
		case choice == 0:
			logger.Info("Пользователь вышел из казино")
			io.WriteLine(player.StatsString())
			io.WriteLine("Всего доброго!")
			return true, nil
		default:
			gameIndex := choice - 1

			if gameIndex >= 0 && gameIndex < len(gamesList) {
				chosenGame := gamesList[gameIndex]
				gameHasBalance, err := chosenGame.Play(ctx, player, io, logger)
				if err != nil {
					return false, err
				}
				if !gameHasBalance {
					return false, nil
				}
			}
		}
	}
	// Если цикл завершился из-за отсутствия денег.
	return false, nil
}

// Start запускает приложение казино.
//
// Выполняет первоначальную инициализацию: создание списка доступных игр,
// запрос имени и стартового баланса у пользователя, создание объекта игрока,
// а затем запускает главное меню (Menu).
func Start(ctx context.Context, io ui.FullIOProvider, logger logservice.IFullLogger) error {
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
	name, err := io.ReadStringCtx(ctx, prompt, true)
	if err != nil {
		logger.Fatal("Ошибка при чтении имени пользователя: %v", err)
		io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
		return err
	}

	// Max = 0 в ReadFloat означает, что максимальное значение не ограничено
	prompt = "И сколько же у тебя денег? "
	balance, err = io.ReadFloatMinCtx(ctx, prompt, 1)
	if err != nil {
		logger.Fatal("Ошибка при чтении баланса пользователя: %v", err)
		io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
		return err
	}

	player := player.NewPlayer(name, balance)

	io.Writef("Привет, %s! Ты начинаешь с балансом %.2f.\n", player.Name, player.Balance)

	exited, err := Menu(ctx, &player, io, logger, gamesList)
	if err != nil {
		logger.Fatal("Меню завершилось ошибкой: %v", err)
		return err
	}

	if exited {
		logger.Info("Игрок %s вышел из казино сам. Итоговый баланс: %.2f", player.Name, player.Balance)
	} else {
		logger.Warn("Игрок %s обанкротился. Итоговый баланс: %.2f", player.Name, player.Balance)
	}

	return nil
}
