package languages

type Translation interface {
	Langauage
	Id() string
	Gender() string
	Value() string
}
