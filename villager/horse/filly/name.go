package filly

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Filly"
	nameCanadianFrench       string = "Palomina"
	nameDutch                string = ""
	nameFrench               string = "Palomina"
	nameGerman               string = "Beate"
	nameItalian              string = "Filippa"
	nameJapanese             string = "7ごう"
	nameLatinAmericanSpanish string = "Séptima"
	nameKorean               string = "7호"
	nameRussian              string = ""
	nameSpanish              string = "Séptima"
	nameSimplifiedChinese    string = ""
	nameTraditionalChinese   string = ""
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
