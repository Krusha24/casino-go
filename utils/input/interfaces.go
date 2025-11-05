package input

type InputProvider interface {
	ReadInt(prompt string, min, max int) (int, error)
	ReadFloat(prompt string, min, max float64) (float64, error)
	ReadString(prompt string, allowSpace bool) (string, error)
}

type OutputWriter interface {
	OutPutF(prompt string, name string, balance float64)
	OutPutLN(prompt string)
}

type FullIOProvider interface {
	InputProvider
	InputProvider
}
