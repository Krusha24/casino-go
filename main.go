// main.go
package main

import (
	"casino/game"
	"casino/utils/io"
)

func main() {
	var ConsoleIO = io.ConsoleIO{}
	game.Start(&ConsoleIO)
}
