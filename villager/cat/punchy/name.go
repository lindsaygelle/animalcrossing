package punchy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Punchy"
	nameCanadianFrench       string = "Cédric"
	nameDutch                string = "Punchy"
	nameFrench               string = "Cédric"
	nameGerman               string = "Julian"
	nameItalian              string = "Felix"
	nameJapanese             string = "ビンタ"
	nameLatinAmericanSpanish string = "Félix"
	nameKorean               string = "빙티"
	nameRussian              string = "Панчи"
	nameSpanish              string = "Félix"
	nameSimplifiedChinese    string = "尔光"
	nameTraditionalChinese   string = "爾光"
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
