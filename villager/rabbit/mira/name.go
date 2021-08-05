package mira

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Mira"
	nameCanadianFrench       string = "Grisette"
	nameDutch                string = "Mira"
	nameFrench               string = "Grisette"
	nameGerman               string = "Mira"
	nameItalian              string = "Martina"
	nameJapanese             string = "ミラコ"
	nameLatinAmericanSpanish string = "Amparo"
	nameKorean               string = "미랑"
	nameRussian              string = "Мира"
	nameSpanish              string = "Amparo"
	nameSimplifiedChinese    string = "蜜拉"
	nameTraditionalChinese   string = "蜜拉"
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
