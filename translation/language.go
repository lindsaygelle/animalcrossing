package translation

type Language interface {
	Code() string
	Name() string
	Value() string
}
