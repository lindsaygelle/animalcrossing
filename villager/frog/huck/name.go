package huck

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Huck"
	nameCanadianFrench       string = "Bajoue"
	nameDutch                string = "Huck"
	nameFrench               string = "Bajoue"
	nameGerman               string = "Knuth"
	nameItalian              string = "Berto"
	nameJapanese             string = "ストロー"
	nameLatinAmericanSpanish string = "Ricardo"
	nameKorean               string = "스트로"
	nameRussian              string = "Хак"
	nameSpanish              string = "Ricardo"
	nameSimplifiedChinese    string = "吸管"
	nameTraditionalChinese   string = "吸管"
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
