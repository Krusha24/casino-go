package ui

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const maxRetries = 3

// ConsoleIO является конкретной реализацией FullIOProvider.
// Предоставляет методы для чтения ввода и вывода информации в консоль.
type ConsoleIO struct{}

func rangeHint(min, max int) string {
	return fmt.Sprintf("[%d..%d]", min, max)
}

func rangeHintF(min, max float64) string {
	return fmt.Sprintf("[%.2f..%.2f]", min, max)
}

func hasAnySpace(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

// ReadInt запрашивает у пользователя ввод целого числа и проверяет,
// находится ли оно в заданном диапазоне [min, max] включительно.
// Реализует метод интерфейса InputProvider.
func (c *ConsoleIO) ReadIntCtx(ctx context.Context, prompt string, min, max int) (int, error) {
	if max < min {
		return 0, fmt.Errorf("невозможный диапазон: min %d > max %d", min, max)
	}
	reader := bufio.NewReader(os.Stdin)
	retries := 0

	for {
		fmt.Print(prompt)

		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return 0, io.EOF
			}
			return 0, err
		}

		line = strings.TrimSpace(line)

		n, err := strconv.Atoi(line)
		if err != nil || n < min || n > max {
			retries++
			fmt.Printf("Введите число %s.\n", rangeHint(min, max))

			if retries >= maxRetries {
				return 0, fmt.Errorf("слишком много невалидных попыток")
			}
			continue
		}
		return n, nil
	}
}

func (c *ConsoleIO) ReadFloatMinCtx(ctx context.Context, prompt string, min float64) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	retries := 0

	for {
		fmt.Print(prompt)

		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return 0, io.EOF
			}
			return 0, err
		}

		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, ",", ".")

		n, err := strconv.ParseFloat(line, 64)
		if err != nil || n < min {
			retries++
			fmt.Printf("Введите число ≥ %.2f.\n", min)

			if retries >= maxRetries {
				return 0, fmt.Errorf("слишком много невалидных попыток")
			}
			continue
		}
		return n, nil
	}
}

func (c *ConsoleIO) ReadFloatRangeCtx(ctx context.Context, prompt string, min, max float64) (float64, error) {
	if max < min {
		return 0, fmt.Errorf("invalid range: min %.2f > max %.2f", min, max)
	}
	reader := bufio.NewReader(os.Stdin)
	retries := 0

	for {
		fmt.Print(prompt)

		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return 0, io.EOF
			}
			return 0, err
		}

		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, ",", ".")

		n, err := strconv.ParseFloat(line, 64)
		if err != nil || n < min || n > max {
			retries++
			fmt.Printf("Введите число %s.\n", rangeHintF(min, max))

			if retries >= maxRetries {
				return 0, fmt.Errorf("слишком много невалидных попыток")
			}
			continue
		}
		return n, nil
	}
}

// ReadString запрашивает у пользователя ввод строки.
// Проверяет, что строка не пустая. Если allowSpace == false, запрещает пробелы.
// Реализует метод интерфейса InputProvider.
func (c *ConsoleIO) ReadStringCtx(ctx context.Context, prompt string, allowSpace bool) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	retries := 0

	for {
		fmt.Print(prompt)

		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return "", io.EOF
			}
			return "", err
		}

		input := strings.TrimSpace(line)
		if input == "" {
			retries++
			fmt.Println("Ошибка. Строка не может быть пустой")
			if retries >= maxRetries {
				return "", fmt.Errorf("слишком много невалидных попыток")
			}
			continue
		}
		if !allowSpace && hasAnySpace(input) {
			retries++
			fmt.Println("Ошибка. Нельзя использовать пробелы.")
			if retries >= maxRetries {
				return "", fmt.Errorf("слишком много невалидных попыток")
			}
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
