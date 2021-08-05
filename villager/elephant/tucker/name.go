package tucker

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tucker"
	nameCanadianFrench       string = "Barry"
	nameDutch                string = "Tucker"
	nameFrench               string = "Barry"
	nameGerman               string = "Thorsten"
	nameItalian              string = "Sventolo"
	nameJapanese             string = "はじめ"
	nameLatinAmericanSpanish string = "Pirolo"
	nameKorean               string = "맘모"
	nameRussian              string = "Такер"
	nameSpanish              string = "Pirolo"
	nameSimplifiedChinese    string = "阿原"
	nameTraditionalChinese   string = "阿原"
)

var (
	name = map[language.Tag]string{
		language.AmericanEnglish:      nameAmericanEnglish,
		language.CanadianFrench:       nameCanadianFrench,
		language.Dutch:                nameDutch,
		language.French:               nameFrench,
		language.German:               nameGerman,
		language.Italian:              nameItalian,
		language.Japanese:             nameJapanese,
		language.Korean:               nameKorean,
		language.LatinAmericanSpanish: nameLatinAmericanSpanish,
		language.Russian:              nameRussian,
		language.Spanish:              nameSpanish,
		language.SimplifiedChinese:    nameSimplifiedChinese,
		language.TraditionalChinese:   nameTraditionalChinese}
)
