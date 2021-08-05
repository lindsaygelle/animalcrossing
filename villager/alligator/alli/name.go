package alli

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Alli"
	nameCanadianFrench       string = "Allie"
	nameDutch                string = "Alli"
	nameFrench               string = "Allie"
	nameGerman               string = "Ali"
	nameItalian              string = "Alli"
	nameJapanese             string = "クロコ"
	nameLatinAmericanSpanish string = "Codrila"
	nameKorean               string = "크로크"
	nameRussian              string = "Элли"
	nameSpanish              string = "Codrila"
	nameSimplifiedChinese    string = "鳄罗思"
	nameTraditionalChinese   string = "鱷羅思"
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
