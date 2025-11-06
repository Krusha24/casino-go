// main.go
package main

import (
	"casino/game"
	"casino/logservice"
	"casino/utils/io"
)

// main является точкой входа в приложение "Казино".
// Здесь происходит инициализация всех необходимых компонентов (ввод/вывод, логирование)
// и запуск главного цикла игры.
func main() {
	// ConsoleIO - это конкретная реализация интерфейса FullIOProvider для работы с консолью.
	var ConsoleIO = io.ConsoleIO{}

	// logger - единый FullLogger, реализующий как системное, так и игровое логирование.
	var logger = logservice.NewFullLogger()

	// Start запускает основную логику приложения, передавая ему зависимости.
	game.Start(&ConsoleIO, logger)
}
