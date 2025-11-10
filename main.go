// main.go
package main

import (
	"casino/game"
	"casino/logservice"
	"casino/utils/ui"
	"context"
	"os"
	"os/signal"
)

// main является точкой входа в приложение "Казино".
// Здесь происходит инициализация всех необходимых компонентов (ввод/вывод, логирование)
// и запуск главного цикла игры.
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// ConsoleIO - это конкретная реализация интерфейса FullIOProvider для работы с консолью.
	var consoleIO = ui.ConsoleIO{}

	// logger - единый FullLogger, реализующий как системное, так и игровое логирование.
	var logger = logservice.NewFullLogger()

	// Start запускает основную логику приложения, передавая ему зависимости.
	if err := game.Start(ctx, &consoleIO, logger); err != nil {
		os.Exit(1)
	}
}
