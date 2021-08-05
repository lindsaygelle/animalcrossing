package liz

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Liz"
	nameCanadianFrench       string = "groonch"
	nameDutch                string = ""
	nameFrench               string = "groonch"
	nameGerman               string = "grrruuuua"
	nameItalian              string = "grunch"
	nameJapanese             string = "なの"
	nameLatinAmericanSpanish string = "ñec-ñec"
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "ñec-ñec"
	nameSimplifiedChinese    string = "嗯呐"
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
