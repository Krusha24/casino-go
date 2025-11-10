package guessnumber

import (
	"casino/logservice"
	"casino/player"
	"casino/utils"
	"casino/utils/ui"
	"context"
)

// Guessnumber является конкретной реализацией интерфейса game.IGame.
type GuessNumber struct{}

// NewGame создает новый экземпляр игры "Угадайка".
func NewGame() *GuessNumber {
	return &GuessNumber{}
}

// Name возвращает название игры для отображения в меню.
// Реализует метод интерфейса game.IGame.
func (g GuessNumber) Name() string {
	return "Поиграть в угадайку"
}

// Play - основной метод, который запускает игровой процесс "Угадайка".
// Реализует метод интерфейса game.IGame.
// Возвращает true, если игрок решил выйти до потери всего баланса,
// и false, если игрок потерял весь баланс.
func (g GuessNumber) Play(ctx context.Context, p *player.Player, io ui.FullIOProvider, logger logservice.IFullLogger) (bool, error) {
	logger.Info("Игрок %s начал игру 'Угадайка' с балансом %.2f", p.Name, p.Balance)

	io.WriteLine("===================================")
	io.WriteLine("== Игра: Угадай число (Guess Game) ==")
	io.WriteLine("===================================")

	for p.Balance > 0 {
		io.Writef("Ваш текущий баланс: %.2f\n", p.Balance)

		choice, err := io.ReadIntCtx(ctx, "Продолжить игру (1) или Выйти в главное меню (2)? ", 1, 2)
		if err != nil {
			logger.Fatal("Ошибка при чтении ввода пользователя: %v", err)
			io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
			return false, err
		}

		if choice == 2 {
			logger.Info("Игрок %s вышел из игры 'Угадайка', баланс %.2f", p.Name, p.Balance)
			return true, nil // Игрок вышел, баланс > 0, возвращаемся в Menu.
		}

		difficulty, err := io.ReadIntCtx(ctx, "Сколько правильных чисел вы хотите?\n1: коэффициент - 10\n2: коэффициент - 4.5\n3: коэффициент - 2.5\n4: коэффициент - 1.7\n", 1, 4)
		if err != nil {
			logger.Fatal("Ошибка при чтении ввода пользователя: %v", err)
			io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
			return false, err
		}

		var winIndexes = utils.CreateWinIndexes(difficulty)

		p.Bet, err = io.ReadFloatRangeCtx(ctx, "Введите вашу ставку: ", 0.01, p.Balance)
		if err != nil {
			logger.Fatal("Ошибка при чтении ввода пользователя: %v", err)
			io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
			return false, err
		}

		chosenNumber, err := io.ReadIntCtx(ctx, "Выберите выигрышное число число от 1 до 10: ", 1, 10)
		if err != nil {
			logger.Fatal("Ошибка при чтении ввода пользователя: %v", err)
			io.WriteLine("Критическая ошибка ввода. Приложение завершается.")
			return false, err
		}

		var isWinner bool
		for _, value := range winIndexes {
			if value == chosenNumber {
				winningSum := calculateNetWinnings(difficulty, p.Bet)
				isWinner = true
				bet := p.Bet
				p.Win(winningSum)

				logger.LogGame(p.Name, bet, "Win", p.Balance)
				io.Writef("Вы угадали! Ваш баланс - %.2f\n", p.Balance)

				break
			}
		}
		if !isWinner {
			bet := p.Bet
			p.Lose()
			logger.LogGame(p.Name, bet, "Lose", p.Balance)
			io.Writef("Вы проиграли, ваш баланс - %.2f\n", p.Balance)
		}
	}
	io.WriteLine("\n!!! У ВАС ЗАКОНЧИЛИСЬ ДЕНЬГИ. ИГРА ОКОНЧЕНА !!!\n")
	logger.Warn("Игрок %s потерял весь баланс и покинул казино.", p.Name)
	return false, nil
}
