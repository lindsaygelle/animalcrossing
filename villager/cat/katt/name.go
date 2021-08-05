package katt

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Katt"
	nameCanadianFrench       string = "Kat"
	nameDutch                string = "Katt"
	nameFrench               string = "Kat"
	nameGerman               string = "Janine"
	nameItalian              string = "Ofelia"
	nameJapanese             string = "ちょい"
	nameLatinAmericanSpanish string = "Melina"
	nameKorean               string = "쵸이"
	nameRussian              string = "Кэт"
	nameSpanish              string = "Melina"
	nameSimplifiedChinese    string = "巧巧"
	nameTraditionalChinese   string = "巧巧"
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
