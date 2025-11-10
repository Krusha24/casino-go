package game

import (
	"casino/logservice"
	"casino/player"
	"casino/utils/ui"
	"context"
)

// IGame определяет общий контракт для всех игр в казино.
// Это позволяет унифицировать запуск любой игры через один метод.
type IGame interface {
	// Play запускает основной игровой цикл для данной игры.
	// Принимает игрока, I/O провайдер и логгер.
	// Возвращает true, если у игрока остался баланс и он вернулся в главное меню,
	// или false, если игра завершилась из-за нулевого баланса.
	Play(ctx context.Context, player *player.Player, io ui.FullIOProvider, logger logservice.IFullLogger) (bool, error)

	// Name возвращает название игры для отображения в меню.
	Name() string
}
