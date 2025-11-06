package logservice

// IGameLogger определяет интерфейс для записи истории игровых раундов (ставок, результатов и баланса).
type IGameLogger interface {
	// LogGame записывает информацию об одном раунде игры в постоянное хранилище.
	LogGame(name string, bet float64, result string, balance float64) error
}

type ISystemLogger interface {
	// Info выводит информационные сообщения, например, начало или конец игры.
	Info(format string, a ...interface{})
	// Warn выводит предупреждения о некритических ситуациях, например, игрок вышел.
	Warn(format string, a ...interface{})
	// Error выводит сообщения об ошибках, не приводящих к остановке приложения
	Error(format string, a ...interface{})
	// Fatal выводит критическую ошибку и завершает работу приложения.
	Fatal(format string, a ...interface{})
}

// IFullLogger объединяет оба контракта: запись истории игр и системное логирование.
// Этот интерфейс используется в качестве зависимости в игровых модулях.
type IFullLogger interface {
	IGameLogger
	ISystemLogger
}
