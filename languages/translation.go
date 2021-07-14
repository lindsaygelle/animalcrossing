package languages

type Translation interface {
	Langauage
	Gender() string
	Value() string
}
