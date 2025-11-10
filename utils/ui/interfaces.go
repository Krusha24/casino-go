package ui

import "context"

// InputProvider определяет методы для чтения данных от пользователя.
type InputProvider interface {
	// ReadInt запрашивает целое число в диапазоне [min, max] включительно.
	ReadIntCtx(ctx context.Context, prompt string, min, max int) (int, error)
	// ReadFloat запрашивает дробное число (float64) в диапазоне [min, max] включительно.
	ReadFloatRangeCtx(ctx context.Context, prompt string, min, max float64) (float64, error)

	ReadFloatMinCtx(ctx context.Context, prompt string, min float64) (float64, error)

	// ReadString запрашивает строку. Если allowSpace=false, запрещает пробелы.
	ReadStringCtx(ctx context.Context, prompt string, allowSpace bool) (string, error)
}

// OutputWriter определяет методы для вывода информации пользователю.
type OutputWriter interface {
	// Write выводит данные без перевода строки.
	Write(a ...interface{})
	// WriteLine выводит данные с переводом строки.
	WriteLine(a ...interface{})
	// Writef выводит форматированные данные, используя синтаксис Printf.
	Writef(format string, a ...interface{})
}

// FullIOProvider объединяет все методы ввода и вывода (InputProvider и OutputWriter).
// Это единый контракт, который мы передаем игровым функциям.
type FullIOProvider interface {
	InputProvider
	OutputWriter
}
