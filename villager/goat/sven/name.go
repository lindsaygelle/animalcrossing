package sven

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sven"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "beuh-heu"
	nameGerman               string = "moooh"
	nameItalian              string = "bee eh"
	nameJapanese             string = "とほほ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "buuueeeeh"
	nameSimplifiedChinese    string = "活活"
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
