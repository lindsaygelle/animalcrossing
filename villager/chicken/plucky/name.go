package plucky

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Plucky"
	nameCanadianFrench       string = "Poulette"
	nameDutch                string = "Plucky"
	nameFrench               string = "Poulette"
	nameGerman               string = "Mareile"
	nameItalian              string = "Cocca"
	nameJapanese             string = "パタヤ"
	nameLatinAmericanSpanish string = "Herminia"
	nameKorean               string = "파타야"
	nameRussian              string = "Плаки"
	nameSpanish              string = "Herminia"
	nameSimplifiedChinese    string = "帕塔雅"
	nameTraditionalChinese   string = "帕塔雅"
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
