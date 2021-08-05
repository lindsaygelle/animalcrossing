package chico

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Chico"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "ma croûte"
	nameGerman               string = "griiins"
	nameItalian              string = "caciotta"
	nameJapanese             string = "ナリ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "quesito"
	nameSimplifiedChinese    string = "哗哗"
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
