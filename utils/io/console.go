package io

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// ConsoleIO является конкретной реализацией FullIOProvider.
// Предоставляет методы для чтения ввода и вывода информации в консоль.
type ConsoleIO struct{}

// ReadInt запрашивает у пользователя ввод целого числа и проверяет,
// находится ли оно в заданном диапазоне [min, max] включительно.
// Реализует метод интерфейса InputProvider.
func (c *ConsoleIO) ReadInt(prompt string, min, max int) (int, error) {
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

// ReadFloat запрашивает у пользователя ввод дробного числа (float64) и проверяет,
// находится ли оно в заданном диапазоне [min, max].
// Если max равен 0, используется math.MaxFloat64. Реализует метод интерфейса InputProvider.
func (c *ConsoleIO) ReadFloat(prompt string, min, max float64) (float64, error) {
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

// ReadString запрашивает у пользователя ввод строки.
// Проверяет, что строка не пустая. Если allowSpace == false, запрещает пробелы.
// Реализует метод интерфейса InputProvider.
func (c *ConsoleIO) ReadString(prompt string, allowSpace bool) (string, error) {
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

// Write выводит данные без перевода строки.
// Реализует метод интерфейса OutputWriter.
func (c *ConsoleIO) Write(a ...interface{}) {
	fmt.Print(a...)
}

// WriteLine выводит данные с переводом строки.
// Реализует метод интерфейса OutputWriter.
func (c *ConsoleIO) WriteLine(a ...interface{}) {
	fmt.Println(a...)
}

// Writef выводит форматированные данные, используя синтаксис Printf.
// Реализует метод интерфейса OutputWriter.
func (c *ConsoleIO) Writef(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// Swritef форматирует строку, используя синтаксис Printf, и ВОЗВРАЩАЕТ ее.
// Реализует метод интерфейса OutputWriter.
func (c *ConsoleIO) Swritef(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
