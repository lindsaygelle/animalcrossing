package viche

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Viche"
	nameCanadianFrench       string = "Mara"
	nameDutch                string = ""
	nameFrench               string = "Mara"
	nameGerman               string = "L-Pyon"
	nameItalian              string = "Vicky"
	nameJapanese             string = "みさき"
	nameLatinAmericanSpanish string = "Ardelta"
	nameKorean               string = "미사키"
	nameRussian              string = ""
	nameSpanish              string = "Ardelta"
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
