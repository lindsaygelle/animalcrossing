package cyd

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cyd"
	nameCanadianFrench       string = "Punk"
	nameDutch                string = "Cyd"
	nameFrench               string = "Punk"
	nameGerman               string = "Sid"
	nameItalian              string = "Sid"
	nameJapanese             string = "パンクス"
	nameLatinAmericanSpanish string = "Ramón"
	nameKorean               string = "펑크스"
	nameRussian              string = "Сид"
	nameSpanish              string = "Ramón"
	nameSimplifiedChinese    string = "庞克斯"
	nameTraditionalChinese   string = "龐克斯"
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
