package timbra

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Timbra"
	nameCanadianFrench       string = "Sélène"
	nameDutch                string = "Timbra"
	nameFrench               string = "Sélène"
	nameGerman               string = "Tippsi"
	nameItalian              string = "Ambra"
	nameJapanese             string = "つかさ"
	nameLatinAmericanSpanish string = "Hebra"
	nameKorean               string = "잔디"
	nameRussian              string = "Тимбра"
	nameSpanish              string = "Hebra"
	nameSimplifiedChinese    string = "阿司"
	nameTraditionalChinese   string = "阿司"
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
