package papi

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Papi"
	nameCanadianFrench       string = "Bourrico"
	nameDutch                string = "Papi"
	nameFrench               string = "Bourrico"
	nameGerman               string = "Friedel"
	nameItalian              string = "Pierino"
	nameJapanese             string = "オカッピ"
	nameLatinAmericanSpanish string = "Bayo"
	nameKorean               string = "마사마"
	nameRussian              string = "Папи"
	nameSpanish              string = "Bayo"
	nameSimplifiedChinese    string = "小冈"
	nameTraditionalChinese   string = "小岡"
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
