package chadder

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Chadder"
	nameCanadianFrench       string = "Mozzar"
	nameDutch                string = "Chadder"
	nameFrench               string = "Mozzar"
	nameGerman               string = "Charlie"
	nameItalian              string = "Gruviero"
	nameJapanese             string = "チーズ"
	nameLatinAmericanSpanish string = "Roque"
	nameKorean               string = "치즈"
	nameRussian              string = "Чаддер"
	nameSpanish              string = "Roque"
	nameSimplifiedChinese    string = "起司"
	nameTraditionalChinese   string = "起司"
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
