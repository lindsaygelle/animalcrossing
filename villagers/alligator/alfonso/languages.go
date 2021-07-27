package alfonso

type chineseSimplified struct{}

func (v chineseSimplified) Value() string {
	return "阿泥"
}

type chineseTraditional struct{}

func (v chineseTraditional) Value() string {
	return "阿泥"
}

type dutch struct{}

func (v dutch) Value() string {
	return "Alfonso"
}

type english struct{}

func (v english) Value() string {
	return "Alfonso"
}

type french struct{}

func (v french) Value() string {
	return "Alphonse"
}

type frenchQuebec struct{}

func (v frenchQuebec) Value() string {
	return "Alphonse"
}

type german struct{}

func (v german) Value() string {
	return "Markus"
}

type italian struct{}

func (v italian) Value() string {
	return "Alfonso"
}

type japanese struct{}

func (v japanese) Value() string {
	return "アルベルト"
}

type korean struct{}

func (v korean) Value() string {
	return "알베르트"
}

type polish struct{}

func (v polish) Value() string {
	return (english{}).Value()
}

type portuguese struct{}

func (v portuguese) Value() string {
	return (english{}).Value()
}

type russian struct{}

func (v russian) Value() string {
	return "Альфонсо"
}

type spanish struct{}

func (v spanish) Value() string {
	return "Kaimán"
}

type spanishLatinAmerica struct{}

func (v spanishLatinAmerica) Value() string {
	return "Kaimán"
}
