package gayle

type chineseSimplified struct{}

func (v chineseSimplified) Value() string {
	return "爱莉"
}

type chineseTraditional struct{}

func (v chineseTraditional) Value() string {
	return "愛莉"
}

type dutch struct{}

func (v dutch) Value() string {
	return "Gayle"
}

type english struct{}

func (v english) Value() string {
	return "Gayle"
}

type french struct{}

func (v french) Value() string {
	return "Odile"
}

type frenchQuebec struct{}

func (v frenchQuebec) Value() string {
	return "Odile"
}

type german struct{}

func (v german) Value() string {
	return "Rosa"
}

type italian struct{}

func (v italian) Value() string {
	return "Codrilla"
}

type japanese struct{}

func (v japanese) Value() string {
	return "アリゲッティ"
}

type korean struct{}

func (v korean) Value() string {
	return "앨리"
}

type polish struct{}

func (v polish) Value() string {
	return "Gayle"
}

type portuguese struct{}

func (v portuguese) Value() string {
	return "Gayle"
}

type russian struct{}

func (v russian) Value() string {
	return "Гейл"
}

type spanish struct{}

func (v spanish) Value() string {
	return "Boni"
}

type spanishLatinAmerica struct{}

func (v spanishLatinAmerica) Value() string {
	return "Boni"
}
