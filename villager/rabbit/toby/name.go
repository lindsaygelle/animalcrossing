package toby

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Toby"
	nameCanadianFrench       string = "Toby"
	nameDutch                string = "Toby"
	nameFrench               string = "Toby"
	nameGerman               string = "Toby"
	nameItalian              string = "Toby"
	nameJapanese             string = "トビー"
	nameLatinAmericanSpanish string = "Toby"
	nameKorean               string = "토비"
	nameRussian              string = "Тоби"
	nameSpanish              string = "Toby"
	nameSimplifiedChinese    string = "咚比"
	nameTraditionalChinese   string = "咚比"
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
