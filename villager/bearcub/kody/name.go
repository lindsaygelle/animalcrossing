package kody

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Kody"
	nameCanadianFrench       string = "Bill"
	nameDutch                string = "Kody"
	nameFrench               string = "Bill"
	nameGerman               string = "Bernd"
	nameItalian              string = "Salomone"
	nameJapanese             string = "アイダホ"
	nameLatinAmericanSpanish string = "Orestes"
	nameKorean               string = "아이다호"
	nameRussian              string = "Коди"
	nameSpanish              string = "Orestes"
	nameSimplifiedChinese    string = "艾德豪"
	nameTraditionalChinese   string = "艾德豪"
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
