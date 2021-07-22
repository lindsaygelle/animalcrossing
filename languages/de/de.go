package de

type de struct{}

func (d de) Code() string {
	return "de"
}

func (d de) Local() string {
	return "Deutsch"
}

func (d de) Name() string {
	return "German"
}
