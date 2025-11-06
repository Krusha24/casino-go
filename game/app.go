// game/logic.go
package game

import (
	"casino/game/guessnumber"
	"casino/logservice"
	"casino/player"
	"casino/utils/io"
)

// Menu запускает главное меню казино и цикл выбора игр, пока у игрока есть деньги.
// Функция принимает указатель на объект игрока, провайдер ввода/вывода (io) и логгер.
// Возвращает false, если игрок решил выйти или его баланс иссяк.
func Menu(player *player.Player, io io.FullIOProvider, logger logservice.IFullLogger) bool {
	logger.Info("Функция меню запустилась")
	var prompt string
	for player.Balance > 0 {
		prompt = "Чего вы хотите?\nВыйти из казино - 0\nПоиграть в угадайку - 1\n"
		choice, err := io.ReadInt(prompt, 0, 1)
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
		case choice == 1:
			// GuessNumber возвращает true, если игрок вышел из игры.
			// Мы игнорируем его возвращаемое значение и продолжаем цикл Menu.
			guessnumber.Play(player, io, logger)
		}
	}
	return false
}

// Start запускает приложение казино.
// Функция выполняет инициализацию (запрос имени и стартового баланса)
// и запускает главное меню (Menu).
func Start(io io.FullIOProvider, logger logservice.IFullLogger) {
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
	Menu(&player, io, logger)
}
