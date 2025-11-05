package io

type InputProvider interface {
	ReadInt(prompt string, min, max int) (int, error)
	ReadFloat(prompt string, min, max float64) (float64, error)
	ReadString(prompt string, allowSpace bool) (string, error)
}

type OutputWriter interface {
	Write(a ...interface{})
	WriteLine(a ...interface{})
	Writef(format string, a ...interface{})
}

type FullIOProvider interface {
	InputProvider
	OutputWriter
}
