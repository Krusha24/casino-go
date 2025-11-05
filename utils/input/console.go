package input

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type ConsoleInput struct{}

func (c *ConsoleInput) ReadInt(prompt string, min, max int) (int, error) {
	var value int
	fmt.Print(prompt)
	for {
		_, err := fmt.Scan(&value)
		if err == nil && value >= min && value <= max {
			return value, nil
		}

		if err != nil {
			fmt.Println("Введите целое число.")
		} else {
			if max == 0 {
				fmt.Printf("Ошибка. Введите число %d.\n", min)
			} else {
				fmt.Printf("Ошибка. Введите число от %d до %d влючительно.\n", min, max)
			}
		}

		// Очистка буфера ввода после ошибки
		var discard string
		fmt.Scanln(&discard)
	}
}

func (c *ConsoleInput) ReadFloat(prompt string, min, max float64) (float64, error) {
	if max == 0 {
		max = math.MaxFloat64
	}

	var value float64

	fmt.Print(prompt)
	for {
		_, err := fmt.Scan(&value)
		if err == nil && value > min && value <= max {
			return value, nil
		}

		if err != nil {
			fmt.Println("Введите целое число.")
		} else {

			if max == math.MaxFloat64 {
				fmt.Printf("Ошибка. Введите число от %.2f.\n", min)
			} else {
				fmt.Printf("Ошибка. Введите число от %.2f до %.2f влючительно.\n", min, max)
			}

		}

		// Очистка буфера ввода после ошибки
		var discard string
		fmt.Scanln(&discard)
	}
}

func (c *ConsoleInput) ReadString(prompt string, allowSpace bool) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Ошибка. Строка не может быть пустой")
			continue
		}

		if !allowSpace && strings.Contains(input, " ") {
			fmt.Println("Ошибка. Нельзя использовать пробелы.")
			continue
		}

		return input, nil
	}
}
