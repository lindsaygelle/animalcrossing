package dora

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Dora"
	nameCanadianFrench       string = "Dora"
	nameDutch                string = "Dora"
	nameFrench               string = "Dora"
	nameGerman               string = "Dora"
	nameItalian              string = "Serena"
	nameJapanese             string = "とめ"
	nameLatinAmericanSpanish string = "Dori"
	nameKorean               string = "티미"
	nameRussian              string = "Дора"
	nameSpanish              string = "Dori"
	nameSimplifiedChinese    string = "杜美"
	nameTraditionalChinese   string = "杜美"
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
