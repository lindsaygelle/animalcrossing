package liz

type chineseSimplified struct{}

func (v chineseSimplified) Value() string {
	return "贝蒂"
}

type chineseTraditional struct{}

func (v chineseTraditional) Value() string {
	return "N/A"
}

type dutch struct{}

func (v dutch) Value() string {
	return "N/A"
}

type english struct{}

func (v english) Value() string {
	return "Liz"
}

type french struct{}

func (v french) Value() string {
	return "Lisette"
}

type frenchQuebec struct{}

func (v frenchQuebec) Value() string {
	return "Lisette"
}

type german struct{}

func (v german) Value() string {
	return "Ruth"
}

type italian struct{}

func (v italian) Value() string {
	return "Drilla"
}

type japanese struct{}

func (v japanese) Value() string {
	return "ストロベリー"
}

type korean struct{}

func (v korean) Value() string {
	return "N/A"
}

type polish struct{}

func (v polish) Value() string {
	return "Liz"
}

type portuguese struct{}

func (v portuguese) Value() string {
	return "Liz"
}

type russian struct{}

func (v russian) Value() string {
	return "N/A"
}

type spanish struct{}

func (v spanish) Value() string {
	return "Natalia"
}

type spanishLatinAmerica struct{}

func (v spanishLatinAmerica) Value() string {
	return "Natalia"
}
