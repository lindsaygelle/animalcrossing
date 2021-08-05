package gwen

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Gwen"
	nameCanadianFrench       string = "Gwen"
	nameDutch                string = "Gwen"
	nameFrench               string = "Gwen"
	nameGerman               string = "Judith"
	nameItalian              string = "Gelinda"
	nameJapanese             string = "ポーラ"
	nameLatinAmericanSpanish string = "Fabiola"
	nameKorean               string = "폴라"
	nameRussian              string = "Гвен"
	nameSpanish              string = "Fabiola"
	nameSimplifiedChinese    string = "宝拉"
	nameTraditionalChinese   string = "寶拉"
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
