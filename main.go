// main.go
package main

import (
	"casino/game"
	"casino/utils/input"
)

func main() {
	var input = input.ConsoleInput{}
	game.Start(&input)
}
