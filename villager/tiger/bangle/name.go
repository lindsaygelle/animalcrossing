package bangle

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bangle"
	nameCanadianFrench       string = "Bengale"
	nameDutch                string = "Bangle"
	nameFrench               string = "Bengale"
	nameGerman               string = "Tamara"
	nameItalian              string = "Tigrizia"
	nameJapanese             string = "ルーズ"
	nameLatinAmericanSpanish string = "Felina"
	nameKorean               string = "루주"
	nameRussian              string = "Бэнгл"
	nameSpanish              string = "Felina"
	nameSimplifiedChinese    string = "卢姿"
	nameTraditionalChinese   string = "盧姿"
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
