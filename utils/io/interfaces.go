package io

// InputProvider определяет методы для чтения данных от пользователя.
type InputProvider interface {
	// ReadInt запрашивает целое число в диапазоне [min, max] включительно.
	ReadInt(prompt string, min, max int) (int, error)
	// ReadFloat запрашивает дробное число (float64) в диапазоне [min, max] включительно.
	ReadFloat(prompt string, min, max float64) (float64, error)
	// ReadString запрашивает строку. Если allowSpace=false, запрещает пробелы.
	ReadString(prompt string, allowSpace bool) (string, error)
}

// OutputWriter определяет методы для вывода информации пользователю.
type OutputWriter interface {
	// Write выводит данные без перевода строки.
	Write(a ...interface{})
	// WriteLine выводит данные с переводом строки.
	WriteLine(a ...interface{})
	// Writef выводит форматированные данные, используя синтаксис Printf.
	Writef(format string, a ...interface{})
	// Swritef форматирует данные, используя синтаксис Printf,
	// и возвращает полученную строку.
	Swritef(format string, a ...interface{}) string
}

// FullIOProvider объединяет все методы ввода и вывода (InputProvider и OutputWriter).
// Это единый контракт, который мы передаем игровым функциям.
type FullIOProvider interface {
	InputProvider
	OutputWriter
}
