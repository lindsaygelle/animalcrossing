package rizzo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rizzo"
	nameCanadianFrench       string = "Sourizzi"
	nameDutch                string = "Rizzo"
	nameFrench               string = "Sourizzi"
	nameGerman               string = "Ricky"
	nameItalian              string = "Rizzo"
	nameJapanese             string = "ちょろきち"
	nameLatinAmericanSpanish string = "Ratolón"
	nameKorean               string = "조르쥐"
	nameRussian              string = "Риццо"
	nameSpanish              string = "Ratolón"
	nameSimplifiedChinese    string = "悄俊"
	nameTraditionalChinese   string = "悄俊"
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
